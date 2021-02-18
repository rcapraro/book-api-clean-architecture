package controller

import (
	"book-api/model"
	"book-api/usecase/interactor"
	"net/http"
)

type bookController struct {
	bookInteractor interactor.BookInteractor
}

type BookController interface {
	GetBooks(c HttpContext) error
}

func NewBookController(bi interactor.BookInteractor) BookController {
	return &bookController{bi}
}

func (bc *bookController) GetBooks(c HttpContext) error {
	var b []*model.Book
	var bo []*interactor.BookOutput
	bo, err := bc.bookInteractor.Get(b)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, bo)
}
