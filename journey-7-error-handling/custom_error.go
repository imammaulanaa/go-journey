package main

import (
	"fmt"
)

type NotFoundError struct {
	Resource string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("%s tidak ditemukan", e.Resource)
}

func cariData(nama string) (string, error) {
	if nama != "Golang" {
		return "", NotFoundError{Resource: nama}
	}
	return "Data ditemukan", nil
}

func main() {
	result, err := cariData("Python")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)
}
