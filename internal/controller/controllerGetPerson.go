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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "error parse struct"})
	}

	verify, err := body.VerifyStruct()

	if verify == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	getPerson := models.ReadPerson(body, c)

	return getPerson
}
