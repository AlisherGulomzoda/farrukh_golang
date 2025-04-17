package my_balance

import (
	"github.com/gofiber/fiber/v3"
	"strconv"
)

func (s *service) SomToDiram(somoni string) (diram int, err error) {

	amount, err := strconv.Atoi(somoni)
	if err != nil {
		return 0, fiber.ErrBadRequest
	}

	result := amount * 100
	return result, nil
}
