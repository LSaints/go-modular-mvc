package repository

import (
	user_management "github.com/LSaints/go-modular-mvc/internal/core/user-management"
)

var users = []user_management.User{
	{ID: "1", Name: "Alice", Email: "alice@example.com"},
	{ID: "2", Name: "Bob", Email: "bob@example.com"},
}

func FindAllUsers() []user_management.User {
	return users
}

func SaveUser(user *user_management.User) error {
	users = append(users, *user)
	return nil
}
