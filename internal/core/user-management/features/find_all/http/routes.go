package http

import (
	"github.com/gin-gonic/gin"
)

func LoadRoutes(r *gin.Engine) {
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", GetAllUsers)
	}
}
