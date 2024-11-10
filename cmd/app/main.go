package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	booksmanagement "github.com/LSaints/go-modular-mvc/internal/core/books-management"
	usermanagement "github.com/LSaints/go-modular-mvc/internal/core/user-management"
	loadsrvvars "github.com/LSaints/go-modular-mvc/internal/shared/config/env_vars/load_server_vars"
	"github.com/LSaints/go-modular-mvc/internal/shared/database"
	"github.com/LSaints/go-modular-mvc/internal/shared/http/adapters"
	"github.com/LSaints/go-modular-mvc/internal/shared/http/interfaces"
)

func Init(r interfaces.Router) {
	booksmanagement.Load(r)
	usermanagement.Load(r)
	loadViews(r)
}

func main() {
	err := loadsrvvars.Execute()
	if err != nil {
		fmt.Println(err)
	}

	_, err = database.Load()

	if err != nil {
		panic(err)
	}

	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	staticFilesPath := path.Join(currentDir, "internal", "shared", "web", "static")

	router := adapters.NewGinRouter()
	Init(router)
	router.Static("/static", staticFilesPath)

	if router == nil {
		panic("Falha ao criar GinRouter")
	}

	err = router.Run(fmt.Sprintf(":%s", loadsrvvars.SERVER_PORT))
	if err != nil {
		panic(err)
	}
}

func loadViews(r interfaces.Router) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	// module: user-management

	// feature: find all
	UserfindAllPage := filepath.Join(currentDir, "internal", "core", "user-management", "features", "find_all", "web", "views", "list.html")

	// feature: create
	UsercreatePage := filepath.Join(currentDir, "internal", "core", "user-management", "features", "create", "web", "views", "register.html")
	UsersuccessPage := filepath.Join(currentDir, "internal", "core", "user-management", "features", "create", "web", "views", "success.html")
	UsererrorPage := filepath.Join(currentDir, "internal", "core", "user-management", "features", "create", "web", "views", "error.html")

	// feature: update
	UserupdatePage := filepath.Join(currentDir, "internal", "core", "user-management", "features", "update", "web", "views", "update.html")
	UsererrorUpdatePage := filepath.Join(currentDir, "internal", "core", "user-management", "features", "update", "web", "views", "error.html")
	UsersuccessUpdatePage := filepath.Join(currentDir, "internal", "core", "user-management", "features", "update", "web", "views", "success.html")

	// module: books-management

	// feature: find all
	BooksfindAllPage := filepath.Join(currentDir, "internal", "core", "books-management", "features", "find_all", "web", "views", "listBooks.html")

	// feature: create
	BooksCreate := filepath.Join(currentDir, "internal", "core", "books-management", "features", "create", "web", "views", "booksRegister.html")
	BooksError := filepath.Join(currentDir, "internal", "core", "books-management", "features", "create", "web", "views", "error.html")
	BooksSuccess := filepath.Join(currentDir, "internal", "core", "books-management", "features", "create", "web", "views", "success.html")

	r.LoadHTMLFiles(
		UserfindAllPage,
		UsercreatePage,
		UsersuccessPage,
		UsererrorPage,
		UserupdatePage,
		UsererrorUpdatePage,
		UsersuccessUpdatePage,
		BooksfindAllPage,
		BooksCreate,
		BooksSuccess,
		BooksError,
	)

	return nil
}
