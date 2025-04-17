package main

import (
	greetHandlerModul "farrukh_golang/cmd/hadler/greet"
	"farrukh_golang/cmd/hadler/math"
	"farrukh_golang/cmd/hadler/my_balance"
	greetServiceModul "farrukh_golang/internal/greet"
	mathSrv "farrukh_golang/internal/math"
	my_balance2 "farrukh_golang/internal/my_balance"
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

	//balanceService := service.NewBalanceService()
	//balanceHandler := handler.NewBalanceHandler(balanceService)

	mybalanceService := my_balance2.New()
	myBalansHandler := my_balance.New(mybalanceService)
	app.Get("/myBalance/:amount", myBalansHandler.MyBalance)

	mathService := mathSrv.New()
	mathHandler := math.New(mathService)

	app.Get("/operation/:operation/a/:a/b/:b", mathHandler.Calculate)

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
