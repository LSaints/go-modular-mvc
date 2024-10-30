package use_case

import (
	"github.com/LSaints/go-modular-mvc/internal/core/user-management/find_all/data"
	"github.com/LSaints/go-modular-mvc/internal/core/user-management/find_all/domain"
)

func GetAllUsers() []domain.User {
	return data.FindAllUsers()
}
