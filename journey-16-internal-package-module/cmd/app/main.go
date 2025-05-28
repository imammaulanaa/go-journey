package main

import (
	"fmt"

	"internal-package-module/internal/repository"
	"internal-package-module/internal/service"
	"internal-package-module/pkg/utils"
)

func main() {
	fmt.Println("Hello from internal package!")

	result := service.Add(10, 5)
	fmt.Println("Add Result:", result)

	user := repository.GetUser(42)
	fmt.Println("User:", user)

	fmt.Println("Uppercase:", utils.ToUpper("archeus"))
}
