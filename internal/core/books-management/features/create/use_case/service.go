package usecase

import (
	"errors"

	"github.com/LSaints/go-modular-mvc/internal/core/books-management/domain"
	"github.com/LSaints/go-modular-mvc/internal/core/books-management/features/create/data"
)

func RegisterBook(book domain.Book) (int64, error) {
	if book.Title == "" || book.Pages <= 0 {
		return 0, errors.New("title or pages cannot be null")
	}

	return data.RegisterBook(book)
}
