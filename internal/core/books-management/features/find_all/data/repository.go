package data

import (
	"github.com/LSaints/go-modular-mvc/internal/core/books-management/domain"
	"github.com/LSaints/go-modular-mvc/internal/shared/database"
)

func FindAllBooks() ([]domain.Book, error) {
	db, err := database.Load()
	defer db.Close()

	result, err := db.Query(`SELECT * FROM books`)
	if err != nil {
		return nil, err
	}

	var books []domain.Book

	for result.Next() {
		var book domain.Book

		if err = result.Scan(
			&book.ID,
			&book.Title,
			&book.Pages,
			&book.UserId,
		); err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}
