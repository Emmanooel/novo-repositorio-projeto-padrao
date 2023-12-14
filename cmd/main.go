package main

import (
	"projetoPadrao/internal/controller"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/person", controller.CreatePerson)
	app.Get("/person", controller.GetPerson)
	app.Get("/login", controller.Login)
	app.Get("/ranking", controller.GetRanking)
	app.Post("/score", controller.InsertScore)
	app.Listen(":3000")
}
