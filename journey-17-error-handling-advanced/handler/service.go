package handler

import (
	"fmt"

	"error-handling-advanced/customerror"
	"error-handling-advanced/utils"
)

func ProcessUser(id string) error {
	if id == "" {
		return customerror.NewValidationError("id")
	}

	if id != "123" {
		return utils.WrapError(customerror.NewNotFoundError("user"), "ProcessUser failed")
	}

	fmt.Println("User processed:", id)
	return nil
}
