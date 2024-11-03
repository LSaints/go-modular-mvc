package http

import (
	"fmt"
	"strconv"

	usecase "github.com/LSaints/go-modular-mvc/internal/core/user-management/features/find_id/use_case"
	"github.com/LSaints/go-modular-mvc/internal/shared/http/interfaces"
)

func FindById(ctx interfaces.Context) {
	userId := ctx.Param("id")
	param, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	_, err = usecase.FindById(param)
	if err != nil {
		fmt.Println(err)
		return
	}
}
