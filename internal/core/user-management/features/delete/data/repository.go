package data

import "github.com/LSaints/go-modular-mvc/internal/shared/database"

func DeleteUser(userId uint64) error {
	query := `DELETE FROM users WHERE id = ?`

	db, err := database.Load()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(query, userId)
	if err != nil {
		return err
	}
	return nil
}
