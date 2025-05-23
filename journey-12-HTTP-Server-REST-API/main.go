package main

import (
	"log"
	"net/http"
	"journey12/handler"
)

func main() {
	http.HandleFunc("/users", handler.UsersHandler)
	http.HandleFunc("/users/create", handler.CreateUserHandler)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
