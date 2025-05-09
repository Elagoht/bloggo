package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/mattn/go-sqlite3"
)

// DatabaseError represents a database-specific error
type DatabaseError struct {
	Code    int
	Message string
	Err     error
}

func (e *DatabaseError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *DatabaseError) Unwrap() error {
	return e.Err
}

// AppError represents an application error with HTTP status
type AppError struct {
	Status  int
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

// NewAppError creates a new application error
func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Status:  code,
		Message: message,
		Err:     err,
	}
}

// Common error instances
var (
	ErrBadRequest     = NewAppError(http.StatusBadRequest, "Bad request", nil)
	ErrUnauthorized   = NewAppError(http.StatusUnauthorized, "Unauthorized", nil)
	ErrForbidden      = NewAppError(http.StatusForbidden, "Forbidden", nil)
	ErrNotFound       = NewAppError(http.StatusNotFound, "Resource not found", nil)
	ErrConflict       = NewAppError(http.StatusConflict, "Resource already exists", nil)
	ErrInternalServer = NewAppError(http.StatusInternalServerError, "Internal server error", nil)
)

// Database error messages
const (
	ErrMsgUniqueConstraint = "A resource with this identifier already exists"
	ErrMsgForeignKey       = "Referenced resource does not exist"
	ErrMsgNotNull          = "Required field cannot be empty"
	ErrMsgTooBig           = "Data exceeds maximum allowed size"
)

// MapDatabaseError maps database errors to application errors
func MapDatabaseError(err error) *AppError {
	if err == nil {
		return nil
	}

	// Log the original error for debugging
	log.Printf("Database error: %v", err)

	// Handle SQLite errors
	if sqliteErr, ok := err.(sqlite3.Error); ok {
		// Log the specific SQLite error code for debugging
		log.Printf("SQLite error code: %d, extended code: %d", sqliteErr.Code, sqliteErr.ExtendedCode)

		switch sqliteErr.Code {
		case sqlite3.ErrConstraint:
			switch sqliteErr.ExtendedCode {
			case sqlite3.ErrConstraintUnique:
				return NewAppError(http.StatusConflict, "A resource with this identifier already exists", err)
			case sqlite3.ErrConstraintForeignKey:
				return NewAppError(http.StatusBadRequest, "Referenced resource does not exist", err)
			case sqlite3.ErrConstraintNotNull:
				return NewAppError(http.StatusBadRequest, "Required field cannot be empty", err)
			}
		case sqlite3.ErrNotFound:
			return ErrNotFound
		case sqlite3.ErrTooBig:
			return NewAppError(http.StatusBadRequest, "Data exceeds maximum allowed size", err)
		}
	}

	// Handle standard database errors
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return ErrNotFound
	default:
		// Log the unhandled error type for debugging
		log.Printf("Unhandled error type: %T", err)
		return NewAppError(http.StatusInternalServerError, "Internal server error", err)
	}
}
