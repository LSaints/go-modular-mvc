package main

import (
	"fmt"
	"os"
	"path/filepath"

	userRoutes "github.com/LSaints/go-modular-mvc/internal/core/user-management/routes"
	server_vars "github.com/LSaints/go-modular-mvc/internal/shared/config/env_vars/load_server_vars"
	"github.com/LSaints/go-modular-mvc/internal/shared/database"
	"github.com/gin-gonic/gin"
)

func main() {
	_, err := database.Load()

	err = server_vars.Execute()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	r.LoadHTMLGlob(filepath.Join(currentDir, "internal", "core", "user-management", "views", "templates") + "/*")

	userRoutes.InitRoutes(r)
	r.Run(fmt.Sprintf(":%s", server_vars.SERVER_PORT))
}
