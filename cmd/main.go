package main

import (
	"farrukh_golang/cmd/hadler/math"
	"farrukh_golang/cmd/hadler/my_balance"
	mathSrv "farrukh_golang/internal/math"
	my_balance2 "farrukh_golang/internal/my_balance"
	"github.com/gofiber/fiber/v3"
	"log"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	mybalanceService := my_balance2.New()
	myBalansHandler := my_balance.New(mybalanceService)
	app.Get("/myBalance/:amount", myBalansHandler.MyBalance)

	mathService := mathSrv.New()
	mathHandler := math.New(mathService)

	app.Get("/operation/:operation/a/:a/b/:b", mathHandler.Calculate)

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
