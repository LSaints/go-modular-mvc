package data

import (
	"github.com/LSaints/go-modular-mvc/internal/core/user-management/domain"
	"github.com/LSaints/go-modular-mvc/internal/shared/database"
)

func FindAllUsers() ([]domain.User, error) {
	db, err := database.Load()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	result, err := db.Query(`SELECT * FROM users`)
	if err != nil {
		return nil, err
	}

	var users []domain.User

	for result.Next() {
		var user domain.User

		if err = result.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
