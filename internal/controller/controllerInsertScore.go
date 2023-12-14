package controller

import (
	"projetoPadrao/internal/domain"
	"projetoPadrao/internal/models"

	"github.com/gofiber/fiber/v2"
)

func ControllerInsertScore(c *fiber.Ctx) error {
	person := new(domain.Person)
	err := c.BodyParser(&person)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "[controller insert score] error parse struct"})
	}

	insertScore := models.InsertScore(person, c)

	return insertScore
}
