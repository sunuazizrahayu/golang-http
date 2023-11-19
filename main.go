package main

import (
	"github.com/gofiber/fiber/v2"
	"golang-http/routes"
	"log"
)

func main() {
	app := fiber.New()

	//init route
	routes.RouteApp(app)

	log.Fatal(app.Listen(":3000"))
}
