package controller

import (
	"projetoPadrao/internal/domain"
	"projetoPadrao/internal/models"

	"github.com/gofiber/fiber/v2"
)

func ControllerLogin(c *fiber.Ctx) error {
	body := new(domain.Login)
	bodyReq := c.BodyParser(body)

	if bodyReq != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "[controller login] error parse struct"})
	}

	verify, err := body.VerifyLogin()
	if verify == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	login := models.FindByLogin(body, c)
	return login
}
