package presenter

import (
	"book-api/model"
	"book-api/usecase/output"
	"book-api/usecase/presenter"
	"strconv"
	"strings"
)

type bookPresenter struct {
}

func NewBookPresenter() presenter.BookPresenter {
	return &bookPresenter{}
}

func (u *bookPresenter) PresentBooks(b []*model.Book) []*output.BookOutput {

	var bookOutputs []*output.BookOutput
	for _, book := range b {
		bookOutputs = append(bookOutputs, &output.BookOutput{
			ISBN:      strconv.Itoa(book.ISBN),
			Title:     book.Title,
			Authors:   strings.Join(book.Authors, ", "),
			Year:      strconv.Itoa(book.Year),
			Publisher: book.Publisher,
		})
	}
	return bookOutputs
}
