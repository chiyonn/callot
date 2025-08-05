package errors

import "fmt"

type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

func NewConfigError(msg string) *AppError {
	return &AppError{Code: 1, Message: fmt.Sprintf("Config error: %s", msg)}
}

func NewValidationError(msg string) *AppError {
	return &AppError{Code: 2, Message: fmt.Sprintf("Validation error: %s", msg)}
}

func NewIOError(msg string) *AppError {
	return &AppError{Code: 3, Message: fmt.Sprintf("I/O error: %s", msg)}
}