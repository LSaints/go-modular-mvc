package main

import (
	"fmt"

	usermanagement "github.com/LSaints/go-modular-mvc/internal/core/user-management"
	loadsrvvars "github.com/LSaints/go-modular-mvc/internal/shared/config/env_vars/load_server_vars"
	"github.com/LSaints/go-modular-mvc/internal/shared/database"
	"github.com/LSaints/go-modular-mvc/internal/shared/http/adapters"
	"github.com/LSaints/go-modular-mvc/internal/shared/http/interfaces"
)

func main() {
	err := loadsrvvars.Execute()
	_, err = database.Load()

	if err != nil {
		panic(err)
	}

	router := adapters.NewGinRouter()
	if router == nil {
		panic("Falha ao criar GinRouter")
	}
	loadModules(router)
	err = router.Run(fmt.Sprintf(":%s", loadsrvvars.SERVER_PORT))
	if err != nil {
		panic(err)
	}
}

func loadModules(r interfaces.Router) {
	usermanagement.Load(r)
}
