package data

import "github.com/LSaints/go-modular-mvc/internal/core/user-management/find_all/domain"

var users = []domain.User{
	{ID: "1", Name: "Alice", Email: "alice@example.com"},
	{ID: "2", Name: "Bob", Email: "bob@example.com"},
}

func FindAllUsers() []domain.User {
	return users
}
