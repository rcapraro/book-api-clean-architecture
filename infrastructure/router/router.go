package router

import (
	"book-api/infrastructure/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewRouter(app *fiber.App, c controller.AppController) *fiber.App {
	app.Use(logger.New())

	app.Get("/books", func(ctx *fiber.Ctx) error {
		return c.GetBooks(ctx)
	})

	return app
}
