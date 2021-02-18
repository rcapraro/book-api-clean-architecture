package presenter

import (
	"book-api/model"
	"book-api/usecase/output"
)

type BookPresenter interface {
	PresentBooks(b []*model.Book) []*output.BookOutput
}
