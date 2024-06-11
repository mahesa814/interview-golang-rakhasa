package repositories

import (
	"interview-api/config"
	"interview-api/models"
)

type BookRepository interface {
	GetByID(id int) (models.Book, error)
	GetAll() ([]models.Book, error)
	Create(book models.Book) error
}

type bookRepository struct{}

func (r *bookRepository) GetByID(id string) (models.Book, error) {
	var book models.Book
	err = config.DB.Get(&book, "SELECT * FROM books WHERE id", id)
	return book, err
}

func (r *bookRepository) GetAll() ([]models.Book, error) {
	var books []models.Book
	var err error
	query = "SELECT * FROM books"
	err = config.DB.Select(&books, query)
	return books, err
}

func (r *bookRepository) Create(book models.Book) error {
	_, err := config.DB.NamedExec("INSERT INTO books (title, author) VALUES (:title, :author)", book)
	return err
}
