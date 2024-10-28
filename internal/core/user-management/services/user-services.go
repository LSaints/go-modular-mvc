package services

import (
	user_management "github.com/LSaints/go-modular-mvc/internal/core/user-management"
	"github.com/LSaints/go-modular-mvc/internal/core/user-management/repository"
)

func GetAllUsers() []user_management.User {
	return repository.FindAllUsers()
}

func CreateUser(user *user_management.User) error {
	return repository.SaveUser(user)
}
