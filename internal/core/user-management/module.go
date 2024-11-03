package usermanagement

import (
	"os"
	"path/filepath"

	createRoute "github.com/LSaints/go-modular-mvc/internal/core/user-management/features/create/http"
	deleteUser "github.com/LSaints/go-modular-mvc/internal/core/user-management/features/delete/http"
	findAllRoute "github.com/LSaints/go-modular-mvc/internal/core/user-management/features/find_all/http"
	"github.com/LSaints/go-modular-mvc/internal/shared/http/interfaces"
)

func Load(r interfaces.Router) {
	loadViews(r)
	registerRoutes(r)
}

func registerRoutes(r interfaces.Router) {
	findAllRoute.LoadRoutes(r)
	createRoute.LoadRoutes(r)
	deleteUser.LoadRoutes(r)

}

func loadViews(r interfaces.Router) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	// feature: find all
	findAllPage := filepath.Join(currentDir, "internal", "core", "user-management", "features", "find_all", "web", "views", "list.html")

	// feature: create
	createPage := filepath.Join(currentDir, "internal", "core", "user-management", "features", "create", "web", "views", "register.html")
	successPage := filepath.Join(currentDir, "internal", "core", "user-management", "features", "create", "web", "views", "success.html")
	errorPage := filepath.Join(currentDir, "internal", "core", "user-management", "features", "create", "web", "views", "error.html")

	r.LoadHTMLFiles(findAllPage, createPage, successPage, errorPage)

	return nil
}
