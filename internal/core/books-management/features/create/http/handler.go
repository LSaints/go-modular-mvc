package http

import (
	"fmt"
	"net/http"

	"github.com/LSaints/go-modular-mvc/internal/core/books-management/domain"
	usecase "github.com/LSaints/go-modular-mvc/internal/core/books-management/features/create/use_case"
	"github.com/LSaints/go-modular-mvc/internal/shared/http/interfaces"
)

func RegisterBook(ctx interfaces.Context) {
	var book domain.Book

	if err := ctx.ShouldBind(&book); err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", map[string]string{"error": "data invalid"})
		fmt.Println(err)
		return
	}

	bookID, err := usecase.RegisterBook(book)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", map[string]interface{}{"error": err.Error()})
		fmt.Println(err)
		return
	}

	ctx.HTML(http.StatusCreated, "success.html", map[string]interface{}{"Book_id": bookID})
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
