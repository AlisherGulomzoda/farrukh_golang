package math

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"log"
	"strconv"
)

func (h *handler) Calculate(c fiber.Ctx) error {
	a1 := c.Params("a")
	b1 := c.Params("b")
	operation := c.Params("operation")
	log.Print(operation)

	if operation != "+" && operation != "-" && operation != "*" && operation != "/" {
		return c.Status(fiber.StatusBadRequest).SendString("Ошибочная операция")
	}

	a, errA := strconv.Atoi(a1)
	b, errB := strconv.Atoi(b1)

	if errA != nil || errB != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "ошибка валидации: a и b должны быть целыми числами",
		})
	}

	var result int
	switch operation {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"msg": "ошибка: деление на ноль",
			})
		}
		result = a / b
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": fmt.Sprintf("Результат вычисления - %d %s %d = %d", a, operation, b, result),
	})
}
