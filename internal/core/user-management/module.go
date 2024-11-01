package user_management

import (
	"os"
	"path/filepath"

	createRoute "github.com/LSaints/go-modular-mvc/internal/core/user-management/features/create/http"
	findAllRoute "github.com/LSaints/go-modular-mvc/internal/core/user-management/features/find_all/http"
	"github.com/gin-gonic/gin"
)

func Load(r *gin.Engine) {
	loadViews(r)
	RegisterRoutes(r)
}

func RegisterRoutes(r *gin.Engine) {
	findAllRoute.LoadRoutes(r)
	createRoute.LoadRoutes(r)
}

func loadViews(r *gin.Engine) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	// feature: find all
	findAllPage := filepath.Join(currentDir, "internal", "core", "user-management", "features", "find_all", "views", "list.html")

	// feature: create
	createPage := filepath.Join(currentDir, "internal", "core", "user-management", "features", "create", "views", "register.html")
	successPage := filepath.Join(currentDir, "internal", "core", "user-management", "features", "create", "views", "success.html")
	errorPage := filepath.Join(currentDir, "internal", "core", "user-management", "features", "create", "views", "error.html")

	r.LoadHTMLFiles(findAllPage, createPage, successPage, errorPage)

	return nil
}
