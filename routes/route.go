package routes

import (
	"github.com/gofiber/fiber/v2"
	"golang-http/controllers"
)

func RouteApp(app *fiber.App) {
	app.Get("/", controllers.WelcomeIndex)
}
