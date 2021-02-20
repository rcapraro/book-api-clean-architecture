package controller

import (
	"book-api/usecase/input"
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
	var bo []*output.BookOutput
	bo, err := bc.bookInteractor.FindAll()
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(bo)
}

func (bc *bookController) CreateBook(c *fiber.Ctx) error {
	b := &input.BookInput{}
	if err := c.BodyParser(b); err != nil {
		return err
	}
	validationErrors := b.Validate()
	if len(validationErrors) > 0 {
		return c.Status(http.StatusBadRequest).JSON(nil)
	}
	if err := bc.bookInteractor.Create(b); err != nil {
		return err
	}
	return c.Status(http.StatusCreated).JSON(b)
}
