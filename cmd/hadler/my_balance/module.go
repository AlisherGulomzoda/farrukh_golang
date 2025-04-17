package my_balance

import (
	"farrukh_golang/internal/my_balance"
	"github.com/gofiber/fiber/v3"
)

type Handler interface {
	MyBalance(c fiber.Ctx) error
}

type handler struct {
	myBalanceService my_balance.Service
}

func New(myBalanceService my_balance.Service) Handler {
	return &handler{myBalanceService: myBalanceService}
}
