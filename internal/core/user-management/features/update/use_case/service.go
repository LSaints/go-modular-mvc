package usecase

import (
	"github.com/LSaints/go-modular-mvc/internal/core/user-management/features/update/data"
)

func UpdateUser(userId uint64, name, email, password string) error {
	err := data.UpdateUser(userId, name, email, password)
	if err != nil {
		return err
	}
	return nil
}
