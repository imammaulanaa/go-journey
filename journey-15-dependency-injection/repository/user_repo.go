package repository

import "depedency-injection/model"

type UserRepository interface {
	GetUserByID(id int) (*model.User, error)
}

type UserRepoImpl struct{}

func (r *UserRepoImpl) GetUserByID(id int) (*model.User, error) {
	return &model.User{ID: id, Name: "Archeus"}, nil
}
