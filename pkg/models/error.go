package models

type HTTPError string

const (
	SignupBadRequest               HTTPError = "SIGNUP_BAD_REQUEST"
	JWTBadRequest                  HTTPError = "JWT_BAD_REQUEST"
	LoginUnauthorized              HTTPError = "LOGIN_UNAUTHORIZED"
	Unauthorized                   HTTPError = "UNAUTHORIZED"
	BadRequest                     HTTPError = "BAD_REQUEST"
	InternalServerError            HTTPError = "INTERNAL_SERVER_ERROR"
	UserPhotoUploadError           HTTPError = "USER_PHOTO_UPLOAD_ERROR"
	ShopForbidden                  HTTPError = "SHOP_FORBBIDEN"
	TemporalFolderNotFound         HTTPError = "TEMPORALFOLDER_NOT_FOUND"
	TemporalPhotoNotFound          HTTPError = "TEMPORALPHOTO_NOT_FOUND"
	RecoveryPasswordBadRequest     HTTPError = "RECOVERY_PASSWORD_BAD_REQUEST"
	NoMorePostsFound               HTTPError = "NO_MORE_POSTS_FOUND"
	NotFound                       HTTPError = "NOT_FOUND"
	PostNotAllowComments           HTTPError = "POST_NOT_ALLOW_COMMENTS"
	UserHasAlreadyReviewedThisShop HTTPError = "USER_ALREADY_MAKE_A_REVIEW_OF_THIS_SHOP"
	PaginationError                HTTPError = "PAGINATION_ERROR_PAGE_OR_PAGESIZE_ARE_NULL"
	UserAlreadyLikedPost           HTTPError = "USER_ALREADY_LIKED_THE_POST"
)

type Error struct {
	Message HTTPError `json:"message"`
}
