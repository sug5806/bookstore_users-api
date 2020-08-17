package errors

import "net/http"

type RestError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func NewVadRequestError(message string) *RestError {
	return &RestError{
		Status:  http.StatusBadRequest,
		Message: message,
		Error:   "bad_request",
	}
}
