package httperror

type HTTPErrorMessage string

const (
	InternalServerError HTTPErrorMessage = "Internal Server Error. Please try again later."
)
