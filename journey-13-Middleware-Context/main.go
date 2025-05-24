package main

import (
	"log"
	"net/http"

	"middleware-context/handler"
	"middleware-context/middleware"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/users", middleware.Logging(middleware.Auth(http.HandlerFunc(handler.UsersHandler))))

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
