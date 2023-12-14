package controller

import (
	"projetoPadrao/internal/domain"
	"projetoPadrao/internal/models"

	"github.com/gofiber/fiber/v2"
)

func GetPerson(c *fiber.Ctx) error {
	body := new(domain.Person)
	bodyReq := c.BodyParser(body)

	if bodyReq != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "[controller get person] error parse struct"})
	}

	getPerson := models.ReadPerson(body, c)

	return getPerson
}

func CreatePerson(c *fiber.Ctx) error {
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

func Login(c *fiber.Ctx) error {
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
