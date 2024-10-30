package database

import (
	"database/sql"
	"log"

	"github.com/LSaints/go-modular-mvc/internal/shared/database/configure_database"
)

func Load() (*sql.DB, error) {
	db, err := configure_database.GetDbConn()
	if err != nil {
		log.Fatalf("Failed to get database connection: %s", err)
		return nil, err
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %s", err)
		return nil, err
	}
	return db, nil
}
