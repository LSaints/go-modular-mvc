package main

import (
	"fmt"

	user_management "github.com/LSaints/go-modular-mvc/internal/core/user-management"
	server_vars "github.com/LSaints/go-modular-mvc/internal/shared/config/env_vars/load_server_vars"
	"github.com/LSaints/go-modular-mvc/internal/shared/database"
	"github.com/gin-gonic/gin"
)

func main() {
	err := server_vars.Execute()
	_, err = database.Load()

	if err != nil {
		panic(err)
	}

	r := gin.Default()
	loadModules(r)
	r.Run(fmt.Sprintf(":%s", server_vars.SERVER_PORT))
}

func loadModules(r *gin.Engine) {
	user_management.Load(r)
}
