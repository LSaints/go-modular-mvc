package http

import (
	"fmt"
	"net/http"

	"github.com/LSaints/go-modular-mvc/internal/core/user-management/domain"
	usecase "github.com/LSaints/go-modular-mvc/internal/core/user-management/features/create/use_case"
	"github.com/LSaints/go-modular-mvc/internal/shared/http/interfaces"
)

func RegisterUser(ctx interfaces.Context) {
	var user domain.User

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", map[string]string{"error": "Dados inválidos"})
		fmt.Println(err)
		return
	}

	userID, err := usecase.RegisterUser(user)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", map[string]interface{}{"error": err.Error()})
		fmt.Println(err)
		return
	}

	ctx.HTML(http.StatusCreated, "success.html", map[string]interface{}{"user_id": userID})
}

func RegisterForm(ctx interfaces.Context) {
	ctx.HTML(http.StatusOK, "register.html", nil)
}

func successPage(ctx interfaces.Context) {
	ctx.HTML(http.StatusCreated, "success.html", nil)
}

func errorPage(ctx interfaces.Context) {
	ctx.HTML(http.StatusUnprocessableEntity, "error.html", nil)
}
