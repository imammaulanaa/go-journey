package utils

import (
	"errors"
	"fmt"
)

func WrapError(err error, context string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", context, err)
}

func IsNotFound(err error) bool {
	return errors.As(err, new(*error))
}
