package http

import (
	"net/http"

	usecase "github.com/LSaints/go-modular-mvc/internal/core/books-management/features/find_all/use_case"
	"github.com/LSaints/go-modular-mvc/internal/shared/http/interfaces"
)

func FindAllBooks(ctx interfaces.Context) {
	books, _ := usecase.FindAllBooks()

	ctx.HTML(http.StatusOK, "listBooks.html", map[string]interface{}{
		"books": books,
	})
}
