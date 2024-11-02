package loaddbvars

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DRIVER_DB           string
	MYSQL_DB            string
	MYSQL_ROOT_PASSWORD string
	MYSQL_PORT          string
	MYSQL_USER          string
	MYSQL_PASSWORD      string
	MYSQL_URL           string
)

func Execute() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	DRIVER_DB = os.Getenv("DRIVER_DB")
	MYSQL_DB = os.Getenv("MYSQL_DB")
	MYSQL_ROOT_PASSWORD = os.Getenv("MYSQL_ROOT_PASSWORD")
	MYSQL_PORT = os.Getenv("MYSQL_PORT")
	MYSQL_USER = os.Getenv("MYSQL_USER")
	MYSQL_PASSWORD = os.Getenv("MYSQL_PASSWORD")
	MYSQL_URL = os.Getenv("MYSQL_URL")
}
