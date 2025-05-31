package repository

import (
	"errors"
	"microservice-architecture/internal/domain"
)

type InMemoryUserRepo struct {
	users map[string]*domain.User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		users: map[string]*domain.User{
			"1": {ID: "1", Name: "Archeus"},
		},
	}
}

func (r *InMemoryUserRepo) GetByID(id string) (*domain.User, error) {
	if user, ok := r.users[id]; ok {
		return user, nil
	}
	return nil, errors.New("user not found")
}
