package loadsrvvars

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	SERVER_URL  string
	SERVER_PORT string
)

func Execute() error {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
		return err
	}

	SERVER_URL = os.Getenv("SERVER_URL")
	SERVER_PORT = os.Getenv("SERVER_PORT")
	return nil
}
