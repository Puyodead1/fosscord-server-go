package fielderror

type FieldErrorCode string

const (
	EMAIL_TYPE_INVALID_EMAIL FieldErrorCode = "EMAIL_TYPE_INVALID_EMAIL"
	EMAIL_ALREADY_REGISTERED FieldErrorCode = "EMAIL_ALREADY_REGISTERED"
	BASE_TYPE_MIN_LENGTH     FieldErrorCode = "BASE_TYPE_MIN_LENGTH"
	BASE_TYPE_MAX_LENGTH     FieldErrorCode = "BASE_TYPE_MAX_LENGTH"
	BASE_TYPE_REQUIRED       FieldErrorCode = "BASE_TYPE_REQUIRED"
	INVALID_LOGIN            FieldErrorCode = "INVALID_LOGIN"
)

var ValidationErrors = map[string]FieldErrorCode{
	"email":    EMAIL_TYPE_INVALID_EMAIL,
	"min":      BASE_TYPE_MIN_LENGTH,
	"max":      BASE_TYPE_MAX_LENGTH,
	"required": BASE_TYPE_REQUIRED,
}

// TODO: min and max should be configurable and able to be different for each field
var FieldErrorMessages = map[FieldErrorCode]string{
	EMAIL_TYPE_INVALID_EMAIL: "Not a well formed email address.",
	EMAIL_ALREADY_REGISTERED: "Email is already registered.",
	BASE_TYPE_MIN_LENGTH:     "Must be 6 or more in length.",
	BASE_TYPE_MAX_LENGTH:     "Must be 72 or fewer in length.",
	BASE_TYPE_REQUIRED:       "This field is required",
	INVALID_LOGIN:            "Login or password is invalid.",
}

func (e FieldErrorCode) String() string {
	return string(e)
}

func (e FieldErrorCode) Message() string {
	return FieldErrorMessages[e]
}
