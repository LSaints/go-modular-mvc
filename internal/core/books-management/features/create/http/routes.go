package http

import "github.com/LSaints/go-modular-mvc/internal/shared/http/interfaces"

func LoadRoutes(r interfaces.Router) {
	booksRoutes := r.Group("/books")
	{
		booksRoutes.POST("/books/register", RegisterBook)

		booksRoutes.GET("/books/register", RegisterForm)
		booksRoutes.GET("/books/register/success", successPage)
		booksRoutes.GET("/books/register/error", errorPage)
	}
}
