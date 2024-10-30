package http

import (
	"fmt"
	"net/http"

	"github.com/LSaints/go-modular-mvc/internal/core/user-management/find_all/use_case"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	users := use_case.GetAllUsers()
	fmt.Println("Usuários:", users)
	c.HTML(http.StatusOK, "list.html", gin.H{
		"users": users,
	})
}
