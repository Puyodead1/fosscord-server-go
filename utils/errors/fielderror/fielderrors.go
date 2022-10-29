package fielderror

type FieldErrorMessage string

const (
	EMAIL_TYPE_INVALID_EMAIL FieldErrorMessage = "Not a well formed email address."
	EMAIL_ALREADY_REGISTERED FieldErrorMessage = "Email is already registered."
	BASE_TYPE_MIN_LENGTH     FieldErrorMessage = "Must be 6 or more in length."
	BASE_TYPE_REQUIRED       FieldErrorMessage = "This field is required"
)
