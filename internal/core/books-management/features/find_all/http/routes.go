package http

import (
	"github.com/LSaints/go-modular-mvc/internal/shared/http/interfaces"
)

func LoadRoutes(r interfaces.Router) {
	booksRoutes := r.Group("/books")
	{
		booksRoutes.GET("/books", FindAllBooks)
	}
}
