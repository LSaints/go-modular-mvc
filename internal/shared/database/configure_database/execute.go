package configure_database

import (
	"database/sql"
	"fmt"

	loaddbvars "github.com/LSaints/go-modular-mvc/internal/shared/config/env_vars/load_database_vars"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Driver string
	User   string
	Url    string
	Pass   string
	Port   string
	Name   string
}

func NewDatabase(driver string, user string, url string, pass string, port string, name string) *Database {
	return &Database{
		Driver: driver,
		User:   user,
		Url:    url,
		Pass:   pass,
		Port:   port,
		Name:   name,
	}
}

func GetDbConn() (*sql.DB, error) {
	loaddbvars.Execute()

	db := NewDatabase(
		loaddbvars.DRIVER_DB,
		loaddbvars.MYSQL_USER,
		loaddbvars.MYSQL_URL,
		loaddbvars.MYSQL_PASSWORD,
		loaddbvars.MYSQL_PORT,
		loaddbvars.MYSQL_DB,
	)

	dbConn, err := sql.Open(db.Driver, getConnectionString(db))
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}

func getConnectionString(db *Database) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db.User, db.Pass, db.Url, db.Port, db.Name)
}
