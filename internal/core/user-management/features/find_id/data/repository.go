package data

import (
	"github.com/LSaints/go-modular-mvc/internal/core/user-management/domain"
	"github.com/LSaints/go-modular-mvc/internal/shared/database"
)

func FindById(userId uint64) (domain.User, error) {
	query := `SELECT * FROM users WHERE id = ?`

	db, err := database.Load()
	if err != nil {
		return domain.User{}, err
	}
	defer db.Close()

	var user domain.User
	err = db.QueryRow(query, userId).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
