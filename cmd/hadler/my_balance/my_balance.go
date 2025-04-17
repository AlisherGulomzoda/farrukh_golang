package my_balance

import (
	"errors"
	"github.com/gofiber/fiber/v3"
)

func (h *handler) MyBalance(c fiber.Ctx) error {
	amountStr := c.Params("amount")

	result, err := h.myBalanceService.SomToDiram(amountStr)
	if err != nil {
		if errors.Is(err, fiber.ErrBadRequest) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"msg": err.Error(),
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": "Internal Server Error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": result,
	})
}
