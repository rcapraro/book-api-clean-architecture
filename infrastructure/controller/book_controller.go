package controller

import (
	"book-api/model"
	"book-api/usecase/interactor"
	"book-api/usecase/output"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type bookController struct {
	bookInteractor interactor.BookInteractor
}

type BookController interface {
	GetBooks(c *fiber.Ctx) error
	CreateBook(c *fiber.Ctx) error
}

func NewBookController(bi interactor.BookInteractor) BookController {
	return &bookController{bi}
}

func (bc *bookController) GetBooks(c *fiber.Ctx) error {
	var b []*model.Book
	var bo []*output.BookOutput
	bo, err := bc.bookInteractor.Get(b)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(bo)
}

func (bc *bookController) CreateBook(c *fiber.Ctx) error {
	b := &model.Book{}
	if err := c.BodyParser(b); err != nil {
		return err
	}
	if err := bc.bookInteractor.Create(b); err != nil {
		return err
	}
	return c.Status(http.StatusCreated).JSON(b)
}
