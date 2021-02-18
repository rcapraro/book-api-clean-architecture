package presenter

import (
	"book-api/model"
	"book-api/usecase/interactor"
)

type BookPresenter interface {
	PresentBooks(b []*model.Book) []*interactor.BookOutput
}
