package controllers

import "github.com/gofiber/fiber/v2"

func WelcomeIndex(c *fiber.Ctx) error {
	return c.SendString("Hello World with GoFiber 👋!")
}
