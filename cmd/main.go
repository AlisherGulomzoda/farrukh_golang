package main

import (
	greetHandlerModul "farrukh_golang/cmd/hadler/greet"
	greetServiceModul "farrukh_golang/internal/greet"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	greetService := greetServiceModul.New()
	greetHandler := greetHandlerModul.New(greetService)

	// Define a route for the GET method on the root path '/'
	app.Get("/assalom", greetHandler.Greet)
	app.Get("/name/:name/gender/:gender", greetHandler.Meet)

	app.Get("/operation/:operation/a/:a/b/:b", func(c fiber.Ctx) error {
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
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
