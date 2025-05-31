package test

import (
	"context"
	"testing"

	"microservice-architecture/internal/repository"
	"microservice-architecture/internal/usecase"
)

func TestGetUser(t *testing.T) {
	repo := repository.NewInMemoryUserRepo()
	uc := usecase.NewUserUsecase(repo)

	user, err := uc.GetUserByID(context.Background(), "1")
	if err != nil {
		t.Fatal(err)
	}

	if user.Name != "Archeus" {
		t.Errorf("expected Archeus, got %s", user.Name)
	}
}
