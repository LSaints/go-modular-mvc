package usermanagement

import (
	"os"
	"path/filepath"

	createRoute "github.com/LSaints/go-modular-mvc/internal/core/user-management/features/create/http"
	deleteUser "github.com/LSaints/go-modular-mvc/internal/core/user-management/features/delete/http"
	findAllRoute "github.com/LSaints/go-modular-mvc/internal/core/user-management/features/find_all/http"
	findById "github.com/LSaints/go-modular-mvc/internal/core/user-management/features/find_id/http"
	update "github.com/LSaints/go-modular-mvc/internal/core/user-management/features/update/http"
	"github.com/LSaints/go-modular-mvc/internal/shared/http/interfaces"
)

func Load(r interfaces.Router) {
	loadViews(r)
	registerRoutes(r)
}

func registerRoutes(r interfaces.Router) {
	findAllRoute.LoadRoutes(r)
	createRoute.LoadRoutes(r)
	findById.LoadRoutes(r)
	deleteUser.LoadRoutes(r)
	update.LoadRoutes(r)

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

	// feature: update
	updatePage := filepath.Join(currentDir, "internal", "core", "user-management", "features", "update", "web", "views", "update.html")
	errorUpdatePage := filepath.Join(currentDir, "internal", "core", "user-management", "features", "update", "web", "views", "error.html")
	successUpdatePage := filepath.Join(currentDir, "internal", "core", "user-management", "features", "update", "web", "views", "success.html")

	r.LoadHTMLFiles(findAllPage, createPage, successPage, errorPage, updatePage, errorUpdatePage, successUpdatePage)

	return nil
}
