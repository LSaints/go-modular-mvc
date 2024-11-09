package booksmanagement

import (
	createRoute "github.com/LSaints/go-modular-mvc/internal/core/books-management/features/create/http"
	findAllRoute "github.com/LSaints/go-modular-mvc/internal/core/books-management/features/find_all/http"
	"github.com/LSaints/go-modular-mvc/internal/shared/http/interfaces"
)

func Load(r interfaces.Router) {
	registerRoutes(r)
}

func registerRoutes(r interfaces.Router) {
	findAllRoute.LoadRoutes(r)
	createRoute.LoadRoutes(r)
}
