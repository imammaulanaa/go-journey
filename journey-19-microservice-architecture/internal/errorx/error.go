package errorx

import "fmt"

type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("Code %d: %s", e.Code, e.Message)
}

func NewNotFound(entity string) *AppError {
	return &AppError{Code: 404, Message: fmt.Sprintf("%s not found", entity)}
}
