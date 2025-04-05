package greet

import "github.com/gofiber/fiber/v3"

func (h *handler) Greet(c fiber.Ctx) error {
	// Send a string response to the client
	return c.SendString("Salom aleykum 👋!")
}
