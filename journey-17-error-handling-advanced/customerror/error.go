package customerror

import "fmt"

type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("Code %d: %s", e.Code, e.Message)
}

func NewNotFoundError(entity string) error {
	return &AppError{
		Code:    404,
		Message: fmt.Sprintf("%s not found", entity),
	}
}

func NewValidationError(field string) error {
	return &AppError{
		Code:    400,
		Message: fmt.Sprintf("Invalid %s", field),
	}
}
