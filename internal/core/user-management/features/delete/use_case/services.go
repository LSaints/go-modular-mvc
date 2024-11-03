package usecase

import "github.com/LSaints/go-modular-mvc/internal/core/user-management/features/delete/data"

func DeleteUser(userId uint64) error {
	err := data.DeleteUser(userId)
	if err != nil {
		return err
	}
	return nil
}
