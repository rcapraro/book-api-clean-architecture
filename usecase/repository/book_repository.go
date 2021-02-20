package repository

import "book-api/model"

type BookRepository interface {
	FindAll() ([]*model.Book, error)
	Save(b *model.Book) error
}
