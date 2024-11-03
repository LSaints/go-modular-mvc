package http

import "github.com/LSaints/go-modular-mvc/internal/shared/http/interfaces"

func LoadRoutes(r interfaces.Router) {
	userRoutes := r.Group("/users")
	{
		userRoutes.PUT("/users/:id", UpdateUser)
		userRoutes.GET("/users/update/:id", UpdateForm)
		userRoutes.GET("/users/update/error", ErrorPage)
		userRoutes.GET("/users/update/success", SuccessPage)
	}
}
