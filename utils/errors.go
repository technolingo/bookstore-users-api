package utils

// APIError is a common error interface
type APIError struct {
	Status  uint16 `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
}
