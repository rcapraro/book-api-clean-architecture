package repository

import "book-api/model"

type BookRepository interface {
	FindAll(b []*model.Book) ([]*model.Book, error)
	Save(b *model.Book) error
}
