package usecase

import (
	"github.com/LSaints/go-modular-mvc/internal/core/books-management/domain"
	"github.com/LSaints/go-modular-mvc/internal/core/books-management/features/find_all/data"
)

func FindAllBooks() ([]domain.Book, error) {
	return data.FindAllBooks()
}
