package controller

import (
	"projetoPadrao/internal/domain"
	"projetoPadrao/internal/models"

	"github.com/gofiber/fiber/v2"
)

func ControllerGetPerson(c *fiber.Ctx) error {
	body := new(domain.Person)
	bodyReq := c.BodyParser(body)

	if bodyReq != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "[controller get person] error parse struct"})
	}

	getPerson := models.ReadPerson(body, c)

	return getPerson
}
