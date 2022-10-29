package errors

type HTTPError struct {
	Code    int                     `json:"code"`
	Errors  *map[string]interface{} `json:"errors,omitempty"`
	Message string                  `json:"message"`
}

// TODO: field error struct
