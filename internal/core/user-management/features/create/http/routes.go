package http

import "github.com/gin-gonic/gin"

func LoadRoutes(r *gin.Engine) {
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/register", RegisterUser)

		userRoutes.GET("/register", RegisterForm)
		userRoutes.GET("/register/success", successPage)
		userRoutes.GET("/register/error", errorPage)
	}
}
