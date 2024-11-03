package data

import (
	"github.com/LSaints/go-modular-mvc/internal/shared/database"
)

func UpdateUser(userId uint64, name, email, password string) error {
	query := `UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?`
	db, err := database.Load()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(query, name, email, password, userId)
	if err != nil {
		return err
	}

	return nil
}
