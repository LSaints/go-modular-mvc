package usecase

import (
	"fmt"

	"github.com/LSaints/go-modular-mvc/internal/core/user-management/domain"
	"github.com/LSaints/go-modular-mvc/internal/core/user-management/features/find_id/data"
)

func FindById(userId uint64) (domain.User, error) {
	user, err := data.FindById(userId)
	if err != nil {
		return domain.User{}, nil
	}
	fmt.Println(user)
	return user, nil
}
