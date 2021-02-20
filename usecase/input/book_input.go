package input

import (
	"book-api/model"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type BookInput struct {
	ISBN      string   `json:"isbn" validate:"isbn"`
	Title     string   `json:"title" validate:"required"`
	Authors   []string `json:"authors" validate:"min=1"`
	Year      int      `json:"year" validate:"min=-1000,max=2100"`
	Publisher string   `json:"publisher"`
}

func (b *BookInput) Validate() validator.ValidationErrors {
	validate = validator.New()
	err := validate.Struct(b)
	if err != nil {
		return err.(validator.ValidationErrors)
	}
	return nil
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
