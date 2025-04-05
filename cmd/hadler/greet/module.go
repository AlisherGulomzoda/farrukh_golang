package greet

import (
	"farrukh_golang/internal/greet"
	"github.com/gofiber/fiber/v3"
)

type Handler interface {
	Greet(c fiber.Ctx) error
	Meet(c fiber.Ctx) error
}

type handler struct {
	greetService greet.Service
}

func New(service greet.Service) Handler {
	return &handler{greetService: service}
}
