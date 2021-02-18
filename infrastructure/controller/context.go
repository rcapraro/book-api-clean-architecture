package controller

import "github.com/gofiber/fiber/v2"

type HttpContext interface {
	Status(status int) *fiber.Ctx
	JSON(i interface{}) error
}
