package http

import (
	"github.com/LSaints/go-modular-mvc/internal/shared/http/interfaces"
)

func LoadRoutes(r interfaces.Router) {
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/users/register", RegisterUser)

		userRoutes.GET("/users/register", RegisterForm)
		userRoutes.GET("/users/register/success", successPage)
		userRoutes.GET("/users/register/error", errorPage)
	}
}
