package errors

// APIError is a common error interface
type APIError struct {
	Status  uint16 `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

// NewBadRequestError return a bad request error with a given message
func NewBadRequestError(message string) *APIError {
	return &APIError{
		Status:  400,
		Error:   "BAD_REQUEST",
		Message: message,
	}
}
