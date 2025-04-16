package main

import (
	greetHandlerModul "farrukh_golang/cmd/hadler/greet"
	"farrukh_golang/cmd/hadler/math"
	greetServiceModul "farrukh_golang/internal/greet"
	"github.com/gofiber/fiber/v3"
	"log"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	greetService := greetServiceModul.New()
	greetHandler := greetHandlerModul.New(greetService)

	// Define a route for the GET method on the root path '/'
	app.Get("/assalom", greetHandler.Greet)
	app.Get("/name/:name/gender/:gender", greetHandler.Meet)

	mathHandler := math.New()

	app.Get("/operation/:operation/a/:a/b/:b", mathHandler.Calculate)

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
