package greet

import (
	"github.com/gofiber/fiber/v3"
	"log"
)

func (h *handler) Meet(c fiber.Ctx) error {
	var (
		name   = c.Params("name")
		gender = c.Params("gender")
	)

	msg, err := h.greetService.Meet(name, gender)
	if err != nil {
		log.Println(err)
	}

	return c.SendString(msg)
}
