package main

import (
	greetHandlerModul "farrukh_golang/cmd/hadler/greet"
	greetServiceModul "farrukh_golang/internal/greet"
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

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
