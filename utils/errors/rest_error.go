package errors

import "net/http"

type RestError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Status:  http.StatusBadRequest,
		Message: message,
		Error:   "bad_request",
	}
}
func NewNotFoundError(message string) *RestError {
	return &RestError{
		Status:  http.StatusNotFound,
		Message: message,
		Error:   "not_found",
	}
}
