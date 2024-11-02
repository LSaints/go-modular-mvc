package http

import (
	"net/http"

	"github.com/LSaints/go-modular-mvc/internal/core/user-management/features/find_all/use_case"
	"github.com/LSaints/go-modular-mvc/internal/shared/http/interfaces"
)

func GetAllUsers(ctx interfaces.Context) {
	users, _ := use_case.GetAllUsers()

	ctx.HTML(http.StatusOK, "list.html", map[string]interface{}{
		"users": users,
	})
}
