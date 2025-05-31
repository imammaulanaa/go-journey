package main

import (
	"log"
	"net/http"

	"microservice-architecture/internal/delivery/http"
	"microservice-architecture/internal/logger"
	"microservice-architecture/internal/repository"
	"microservice-architecture/internal/usecase"
)

func main() {
	logger.Init()
	logger.Log.Info("Starting Journey 20 - Clean Arch...")

	repo := repository.NewInMemoryUserRepo()
	uc := usecase.NewUserUsecase(repo)
	handler := http.NewUserHandler(uc)

	mux := http.NewServeMux()
	mux.Handle("/user", http.LoggingMiddleware(http.HandlerFunc(handler.GetUser)))

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
