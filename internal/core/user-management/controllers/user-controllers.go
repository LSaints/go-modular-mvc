package controllers

import (
	"fmt"
	"net/http"

	user_management "github.com/LSaints/go-modular-mvc/internal/core/user-management"
	"github.com/LSaints/go-modular-mvc/internal/core/user-management/services"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	users := services.GetAllUsers()
	fmt.Println("Usuários:", users)
	c.HTML(http.StatusOK, "list.html", gin.H{
		"users": users,
	})
}

func CreateUserForm(c *gin.Context) {
	c.HTML(http.StatusOK, "create.html", nil)
}

func CreateUser(c *gin.Context) {
	var user user_management.User

	if err := c.ShouldBind(&user); err != nil {
		c.HTML(http.StatusBadRequest, "create.html", gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateUser(&user); err != nil {
		c.HTML(http.StatusInternalServerError, "create.html", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, "/users/success")
}

func SuccessPage(c *gin.Context) {
	c.HTML(http.StatusOK, "success.html", nil)
}
