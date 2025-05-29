package main

import (
	"fmt"
	"log"

	"error-handling-advanced/handler"
)

func main() {
	err := handler.ProcessUser("")
	if err != nil {
		log.Println("Error:", err)
	}

	err = handler.ProcessUser("456")
	if err != nil {
		log.Println("Error:", err)
	}

	err = handler.ProcessUser("123")
	if err != nil {
		log.Println("Error:", err)
	} else {
		fmt.Println("Operation successful!")
	}
}
