package service

import (
	"testing"
	"depedency-injection/repository"
)

func TestGetUserInfo(t *testing.T) {
	mockRepo := &repository.UserRepoMock{}
	service := NewUserService(mockRepo)

	service.GetUserInfo(1) // Harus mencetak User dengan nama "Mock User"
}
