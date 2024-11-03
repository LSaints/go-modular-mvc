package http

import "github.com/LSaints/go-modular-mvc/internal/shared/http/interfaces"

func LoadRoutes(r interfaces.Router) {
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/users/:id", FindById)
	}
}
