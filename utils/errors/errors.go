package fcerrors

type HTTPError struct {
	Code    int                    `json:"code"`
	Errors  *map[string]FieldError `json:"errors,omitempty"`
	Message string                 `json:"message"`
}

// _errors
type FieldErrorErrors struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
type FieldError struct {
	EErrors []FieldErrorErrors `json:"_errors"`
}
