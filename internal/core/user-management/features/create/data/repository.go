package data

import (
	"github.com/LSaints/go-modular-mvc/internal/core/user-management/domain"
	"github.com/LSaints/go-modular-mvc/internal/shared/database"
)

func RegisterUser(user domain.User) (int64, error) {
	query := `INSERT INTO users (name, email, password) VALUES (?, ?, ?)`

	db, err := database.Load()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	result, err := db.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return userID, nil
}
