package middleware

import (
	"fmt"
	"github.com/SantaPesca/baselib/pkg/models"
	"github.com/SantaPesca/baselib/pkg/repository"
	"github.com/SantaPesca/baselib/pkg/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"github.com/lib/pq"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"net/http"
)

type Middleware struct{}

/* 		Función encargada de verificar que el usuario tiene permisos
para acceder al controller y que el JWT es correcto
*/

func (m Middleware) MiddleWare(next http.HandlerFunc, db *gorm.DB, rdb redis.Cmdable, action string, subject string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var e models.Error
		redisRepo := repository.RedisRepository{}

		authHeader := request.Header.Get("Authorization")
		bearerString, bearerToken := utils.GetToken(authHeader)

		if authHeader == "" || bearerString != "Bearer" || bearerToken == "" {
			e.Message = models.JWTBadRequest
			utils.MyLog.Println("Error in header (authHeader or bearerToken problem)")
			utils.RespondWithError(writer, http.StatusUnauthorized, e)
			return
		} else {
			token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there was an error")
				}
				return []byte(viper.GetString("secret.jwt")), nil
			})

			if err != nil {
				e.Message = models.Unauthorized
				utils.MyLog.Println("An error occurred signing the token: ", err)
				utils.RespondWithError(writer, http.StatusUnauthorized, e)
				return
			}

			if token != nil {
				if token.Valid && CheckPermissions(db, token, action, subject) {
					getErr := redisRepo.CheckIfTokenExists(rdb, bearerToken)
					if getErr == redis.Nil {
						e.Message = models.Unauthorized
						utils.RespondWithError(writer, http.StatusUnauthorized, e)
						utils.MyLog.Println("Bearer token not exists in redis: ", getErr)
						return
					} else if getErr != nil {
						e.Message = models.InternalServerError
						utils.RespondWithError(writer, http.StatusInternalServerError, e)
						utils.MyLog.Println("Error checking token in redis: ", getErr)
						return
					} else {
						next.ServeHTTP(writer, request)
					}
				} else {
					e.Message = models.Unauthorized
					utils.RespondWithError(writer, http.StatusUnauthorized, e)
					return
				}
			}
		}
	}
}

func CheckPermissions(db *gorm.DB, token *jwt.Token, action string, subject string) bool {
	// obtener el email del usuario del token
	claims := token.Claims.(jwt.MapClaims)
	email := claims["email"]

	// obtener los roles del usuario con este email
	postgresRepo := repository.PostgresRepository{}
	actions, subjects, err := postgresRepo.GetUserActionAndSubjectByEmail(db, email.(string))
	if err != nil {
		utils.MyLog.Println("Cannot get user roles: ", err)
		return false
	}

	// chequear si la action y el subject esta dentro del rol
	if CheckAction(action, actions) && CheckSubject(subject, subjects) {
		return true
	} else {
		return false
	}
}

func CheckAction(action string, roleActions pq.StringArray) bool {
	for _, roleAction := range roleActions {
		if roleAction == action || roleAction == "manage" {
			return true
		}
	}
	return false
}

func CheckSubject(subject string, roleSubjects pq.StringArray) bool {
	for _, roleSubject := range roleSubjects {
		if roleSubject == subject || roleSubject == "all" {
			return true
		}
	}
	return false
}
