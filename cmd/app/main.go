package main

import (
	"os"
	"path/filepath"

	userRoutes "github.com/LSaints/go-modular-mvc/internal/core/user-management/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	r.LoadHTMLGlob(filepath.Join(currentDir, "internal", "core", "user-management", "views", "templates") + "/*")

	userRoutes.InitRoutes(r)
	r.Run(":8080")
}
