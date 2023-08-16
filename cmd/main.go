package main

import (
	"projetoPadrao/internal/controller"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/person", controller.ControllerRegisterPerson)
	app.Get("/person", controller.ControllerGetPerson)
	app.Listen(":3000")
}
