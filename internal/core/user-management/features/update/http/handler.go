package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/LSaints/go-modular-mvc/internal/core/user-management/features/find_id/data"
	usecase "github.com/LSaints/go-modular-mvc/internal/core/user-management/features/update/use_case"
	"github.com/LSaints/go-modular-mvc/internal/shared/http/interfaces"
)

func UpdateUser(ctx interfaces.Context) {
	userId := ctx.Param("id")

	name := ctx.Query("name")
	email := ctx.Query("email")
	password := ctx.Query("password")

	param, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", "Invalid user ID")
		return
	}

	err = usecase.UpdateUser(param, name, email, password)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "success.html", nil)
}

func UpdateForm(ctx interfaces.Context) {
	userId := ctx.Param("id")
	param, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		fmt.Println(err)
		ctx.HTML(http.StatusBadRequest, "error.html", nil)
		return
	}

	user, err := data.FindById(param)
	if err != nil {
		fmt.Println(err)
		ctx.HTML(http.StatusNotFound, "error.html", nil)
		return
	}

	ctx.HTML(http.StatusOK, "update.html", user)
}

func ErrorPage(ctx interfaces.Context) {
	ctx.HTML(http.StatusBadRequest, "error.html", nil)
}

func SuccessPage(ctx interfaces.Context) {
	ctx.HTML(http.StatusBadRequest, "success.html", nil)
}
