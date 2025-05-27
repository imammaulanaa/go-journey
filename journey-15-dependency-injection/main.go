package main

import (
	"depedency-injection/repository"
	"depedency-injection/service"
)

func main() {
	repo := &repository.UserRepoImpl{}
	userService := service.NewUserService(repo)

	userService.GetUserInfo(42)
}
