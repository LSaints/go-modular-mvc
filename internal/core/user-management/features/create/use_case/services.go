package usecase

import (
	"errors"

	"github.com/LSaints/go-modular-mvc/internal/core/user-management/domain"
	"github.com/LSaints/go-modular-mvc/internal/core/user-management/features/create/data"
)

func RegisterUser(user domain.User) (int64, error) {
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return 0, errors.New("name, email or password cannot be null")
	}

	return data.RegisterUser(user)
}
