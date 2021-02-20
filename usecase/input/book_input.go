package input

import (
	"book-api/model"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type BookInput struct {
	ISBN      string   `validate:"isbn"`
	Title     string   `validate:"required"`
	Authors   []string `validate:"min=1"`
	Year      int      `validate:"min=-1000,max=2100"`
	Publisher string
}

func (b *BookInput) Validate() validator.ValidationErrors {
	err := validate.Struct(b)
	return err.(validator.ValidationErrors)
}

func (b *BookInput) ToModel() *model.Book {
	return &model.Book{
		ISBN:      b.ISBN,
		Title:     b.Title,
		Authors:   b.Authors,
		Year:      b.Year,
		Publisher: b.Publisher,
	}
}
