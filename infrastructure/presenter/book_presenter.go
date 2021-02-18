package presenter

import (
	"book-api/model"
	"book-api/usecase/interactor"
	"book-api/usecase/presenter"
	"strconv"
	"strings"
)

type bookPresenter struct {
}

func NewUserPresenter() presenter.BookPresenter {
	return &bookPresenter{}
}

func (u bookPresenter) PresentBooks(b []*model.Book) []*interactor.BookOutput {

	var bookOutputs []*interactor.BookOutput
	for _,book := range b {
		bookOutputs = append(bookOutputs, &interactor.BookOutput{
			ISBN:      strconv.FormatUint(uint64(book.ISBN), 10),
			Title:     book.Title,
			Authors:   strings.Join(book.Authors, ", "),
			Year:      strconv.FormatUint(uint64(book.Year), 10),
			Publisher: book.Publisher,
		})
	}
	return bookOutputs
}

