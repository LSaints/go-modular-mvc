package http

import (
	"fmt"
	"net/http"

	"github.com/LSaints/go-modular-mvc/internal/core/user-management/domain"
	"github.com/LSaints/go-modular-mvc/internal/core/user-management/features/create/use_case"
	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var user domain.User

	if err := c.ShouldBind(&user); err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": "Dados inválidos"})
		fmt.Println(err)
		return
	}

	userID, err := use_case.RegisterUser(user)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	c.HTML(http.StatusCreated, "success.html", gin.H{"user_id": userID})
}

func RegisterForm(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func successPage(c *gin.Context) {
	c.HTML(http.StatusCreated, "success.html", nil)
}

func errorPage(c *gin.Context) {
	c.HTML(http.StatusUnprocessableEntity, "error.html", nil)
}
