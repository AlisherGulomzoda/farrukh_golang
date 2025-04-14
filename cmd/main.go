package main

import (
	greetHandlerModul "farrukh_golang/cmd/hadler/greet"
	greetServiceModul "farrukh_golang/internal/greet"
	"fmt"
	"log"

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
		a := c.Params("a")
		b := c.Params("b")
		operation := c.Params("operation")
		msg := fmt.Sprintf("operation - %s\n a - %s\n b - %s", operation, a, b)
		if len(operation) == 0 {
			return c.Status(fiber.StatusBadRequest).SendString("Ошибочная операция")

		}

		return c.SendString(msg)

	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
