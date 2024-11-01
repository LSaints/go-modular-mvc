package http

import (
	"net/http"

	"github.com/LSaints/go-modular-mvc/internal/core/user-management/features/find_all/use_case"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	users, _ := use_case.GetAllUsers()

	c.HTML(http.StatusOK, "list.html", gin.H{
		"users": users,
	})
}
