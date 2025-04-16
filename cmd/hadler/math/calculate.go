package math

import (
	"errors"
	"github.com/gofiber/fiber/v3"
)

func (h *handler) Calculate(c fiber.Ctx) error {
	a := c.Params("a")
	b := c.Params("b")
	operation := c.Params("operation")

	msg, err := h.mathService.Calculate(operation, a, b)
	if err != nil {
		if errors.Is(err, fiber.ErrBadRequest) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"msg": msg,
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": "Internal Server Error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": msg,
	})
}
