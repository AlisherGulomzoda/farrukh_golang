package math

import (
	"farrukh_golang/internal/math"
	"github.com/gofiber/fiber/v3"
)

type Handler interface {
	Calculate(c fiber.Ctx) error
}

type handler struct {
	mathService math.Service
}

func New(mathService math.Service) Handler {
	return &handler{
		mathService: mathService,
	}
}
