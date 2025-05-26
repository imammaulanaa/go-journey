package main

import (
	"log"
	"net/http"

	"database-access/config"
	"database-access/handler"
)

func main() {
	config.InitDB()

	http.HandleFunc("/users", handler.GetUsersHandler)
	http.HandleFunc("/users/create", handler.CreateUserHandler)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
