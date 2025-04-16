package math

import "github.com/gofiber/fiber/v3"

type Handler interface {
	Calculate(c fiber.Ctx) error
}

type handler struct{}

func New() Handler {
	return &handler{}
}
