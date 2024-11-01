package use_case

import (
	"github.com/LSaints/go-modular-mvc/internal/core/user-management/domain"
	"github.com/LSaints/go-modular-mvc/internal/core/user-management/features/find_all/data"
)

func GetAllUsers() ([]domain.User, error) {
	return data.FindAllUsers()
}
