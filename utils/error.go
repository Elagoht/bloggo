package utils

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/mattn/go-sqlite3"
)

type AppError struct {
	Status  int
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Status:  code,
		Message: message,
		Err:     err,
	}
}

var (
	ErrBadRequest     = NewAppError(http.StatusBadRequest, "Bad request", nil)
	ErrUnauthorized   = NewAppError(http.StatusUnauthorized, "Unauthorized", nil)
	ErrForbidden      = NewAppError(http.StatusForbidden, "Forbidden", nil)
	ErrNotFound       = NewAppError(http.StatusNotFound, "Resource not found", nil)
	ErrConflict       = NewAppError(http.StatusConflict, "Conflict error", nil)
	ErrInternalServer = NewAppError(http.StatusInternalServerError, "Internal server error", nil)
)

func MapDatabaseError(err error) *AppError {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return ErrNotFound
	case errors.Is(err, sqlite3.ErrConstraintUnique):
		return ErrConflict
	case errors.Is(err, sqlite3.ErrConstraintForeignKey):
		return ErrBadRequest
	case errors.Is(err, sqlite3.ErrConstraintNotNull):
		return ErrBadRequest
	case errors.Is(err, sqlite3.ErrNotFound):
		return ErrNotFound
	case errors.Is(err, sqlite3.ErrTooBig):
		return ErrBadRequest
	case errors.Is(err, sqlite3.ErrConstraintNotNull):
		return ErrBadRequest
	default:
		return NewAppError(http.StatusInternalServerError, "Internal server error", err)
	}
}
