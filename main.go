package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"golang-http/routes"
	"log"
	"os"
)

func main() {
	//load config
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	//gofiber app
	app := fiber.New()

	//init route
	routes.RouteApp(app)

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "3000"
	}
	log.Fatal(app.Listen(":" + appPort))
}
