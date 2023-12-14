package main

import (
	"projetoPadrao/internal/controller"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/person", controller.ControllerRegisterPerson)
	app.Get("/person", controller.ControllerGetPerson)
	app.Get("/login", controller.ControllerLogin)
	app.Get("/ranking", controller.ControllerGetRanking)
	app.Post("/score", controller.ControllerInsertScore)
	app.Listen(":3000")
}
