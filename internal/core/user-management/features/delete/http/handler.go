package http

import (
	"fmt"
	"strconv"

	usecase "github.com/LSaints/go-modular-mvc/internal/core/user-management/features/delete/use_case"
	"github.com/LSaints/go-modular-mvc/internal/shared/http/interfaces"
)

func DeleteUser(ctx interfaces.Context) {
	userId := ctx.Param("id")
	param, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	err = usecase.DeleteUser(param)
	if err != nil {
		fmt.Println(err)
		return
	}
}
