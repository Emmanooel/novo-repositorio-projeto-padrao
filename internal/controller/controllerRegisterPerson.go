package controller

import (
	"projetoPadrao/internal/domain"
	"projetoPadrao/internal/models"

	"github.com/gofiber/fiber/v2"
)

func ControllerRegisterPerson(c *fiber.Ctx) error {
	body := new(domain.Person)
	bodyReq := c.BodyParser(body)

	if bodyReq != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "[controller register person] error parse struct"})
	}

	filledStruct, err := body.VerifyPerson()

	if filledStruct == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "[controller register person]" + err})
	}

	createPerson := models.CreatePerson(body, c)

	return createPerson
}
