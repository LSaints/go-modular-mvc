package user_management

import (
	"os"
	"path/filepath"

	"github.com/LSaints/go-modular-mvc/internal/core/user-management/find_all/http"
	"github.com/gin-gonic/gin"
)

func Load(r *gin.Engine) {
	loadView(r, "user-management", "find_all")
	loadRoutes(r)
}

func loadRoutes(r *gin.Engine) {
	http.InitRoutes(r)
}

func loadView(r *gin.Engine, moduleName string, feature string) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	r.LoadHTMLGlob(filepath.Join(currentDir, "internal", "core", moduleName, feature, "views") + "/*")
	return nil
}
