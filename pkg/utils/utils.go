package utils

import (
	"encoding/json"
	"github.com/SantaPesca/baselib/pkg/models"
	guuid "github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"strings"
)

var MyLog = log.New(os.Stderr, "[ERROR] ", log.LstdFlags|log.Lshortfile)

func RespondWithError(writer http.ResponseWriter, status int, error models.Error) {
	writer.WriteHeader(status)
	err := json.NewEncoder(writer).Encode(error)
	if err != nil {
		MyLog.Fatal("Cannot encode error: ", err)
	}
}

func RespondJSON(writer http.ResponseWriter, data interface{}) {
	err := json.NewEncoder(writer).Encode(data)
	if err != nil {
		MyLog.Fatal("Cannot encode data: ", err)
	}
}

func GenUUID() guuid.UUID {
	id := guuid.New()
	return id
}

func GetToken(authHeader string) (string, string) {
	bearerToken := strings.Split(authHeader, " ")

	if len(bearerToken) == 1 {
		return bearerToken[0], ""
	}

	return bearerToken[0], bearerToken[1]
}

func HashPassword(password string) (string, error) {
	//Encriptado de password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return "", err
	}

	//Se parsea la password a String y se retorna
	return string(hash), err
}
