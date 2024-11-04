package usermanagement

import (
	createRoute "github.com/LSaints/go-modular-mvc/internal/core/user-management/features/create/http"
	deleteUser "github.com/LSaints/go-modular-mvc/internal/core/user-management/features/delete/http"
	findAllRoute "github.com/LSaints/go-modular-mvc/internal/core/user-management/features/find_all/http"
	findById "github.com/LSaints/go-modular-mvc/internal/core/user-management/features/find_id/http"
	update "github.com/LSaints/go-modular-mvc/internal/core/user-management/features/update/http"
	"github.com/LSaints/go-modular-mvc/internal/shared/http/interfaces"
)

func Load(r interfaces.Router) {
	registerRoutes(r)
}

func registerRoutes(r interfaces.Router) {
	findAllRoute.LoadRoutes(r)
	createRoute.LoadRoutes(r)
	findById.LoadRoutes(r)
	deleteUser.LoadRoutes(r)
	update.LoadRoutes(r)

}
