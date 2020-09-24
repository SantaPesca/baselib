package models

type HTTPError string

const (
	SignupBadRequest           HTTPError = "SIGNUP_BAD_REQUEST"
	JWTBadRequest              HTTPError = "JWT_BAD_REQUEST"
	LoginUnauthorized          HTTPError = "LOGIN_UNAUTHORIZED"
	Unauthorized               HTTPError = "UNAUTHORIZED"
	BadRequest                 HTTPError = "BAD_REQUEST"
	InternalServerError        HTTPError = "INTERNAL_SERVER_ERROR"
	UserPhotoUploadError       HTTPError = "USER_PHOTO_UPLOAD_ERROR"
	ShopForbidden              HTTPError = "SHOP_FORBBIDEN"
	TemporalFolderNotFound     HTTPError = "TEMPORALFOLDER_NOT_FOUND"
	TemporalPhotoNotFound      HTTPError = "TEMPORALPHOTO_NOT_FOUND"
	RecoveryPasswordBadRequest HTTPError = "RECOVERY_PASSWORD_BAD_REQUEST"
	NoMorePostsFound           HTTPError = "NO_MORE_POSTS_FOUND"
)

type Error struct {
	Message HTTPError `json:"message"`
}
