package data

import (
	"github.com/LSaints/go-modular-mvc/internal/core/books-management/domain"
	"github.com/LSaints/go-modular-mvc/internal/shared/database"
)

func RegisterBook(book domain.Book) (int64, error) {
	query := `INSERT INTO books (title, pages, userId) VALUES (?, ?, ?)`

	db, err := database.Load()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	result, err := db.Exec(query, book.Title, book.Pages, book.UserId)
	if err != nil {
		return 0, err
	}

	bookID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return bookID, nil
}
