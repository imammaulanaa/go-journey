package repository

import "depedency-injection/model"

type UserRepoMock struct{}

func (r *UserRepoMock) GetUserByID(id int) (*model.User, error) {
	return &model.User{ID: id, Name: "Mock User"}, nil
}

