package math

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

func (s *service) Calculate(operation, a1, b1 string) (msg string, err error) {

	if operation != "+" && operation != "-" && operation != "*" && operation != "/" {
		return "Ошибочная операция", fiber.ErrBadRequest
	}

	a, errA := strconv.Atoi(a1)
	b, errB := strconv.Atoi(b1)

	if errA != nil || errB != nil {
		return "ошибка валидации: a и b должны быть целыми числами", fiber.ErrBadRequest
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
			return "ошибка: деление на ноль", fiber.ErrBadRequest
		}
		result = a / b
	}

	return fmt.Sprintf("Результат вычисления - %d %s %d = %d", a, operation, b, result), nil
}
