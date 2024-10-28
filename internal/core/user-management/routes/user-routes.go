package routes

import (
	"github.com/LSaints/go-modular-mvc/internal/core/user-management/controllers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", controllers.GetAllUsers)
		userRoutes.GET("/create", controllers.CreateUserForm)
		userRoutes.POST("/create", controllers.CreateUser)
		userRoutes.GET("/success", controllers.SuccessPage)
	}
}
