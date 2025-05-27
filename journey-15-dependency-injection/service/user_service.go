package service

import (
	"fmt"
	"depedency-injection/model"
	"depedency-injection/repository"
)

type UserService struct {
	Repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetUserInfo(id int) {
	user, err := s.Repo.GetUserByID(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("User: %+v\n", user)
}
