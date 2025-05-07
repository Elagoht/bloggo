package utils

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
}

func (error *Error) Error() string {
	return error.Message
}

func NewError(message string) *Error {
	return &Error{Message: message}
}

func NewErrorWithStatus(error *Error, writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNotFound)
	json.NewEncoder(writer).Encode(error)
}

func HandleError(err error, writer http.ResponseWriter) {
	switch err := err.(type) {
	case *Error:
		NewErrorWithStatus(err, writer)
	default:
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
