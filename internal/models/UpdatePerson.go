package models

import (
	"projetoPadrao/internal/domain"
	"projetoPadrao/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func InsertScore(person *domain.Person, c *fiber.Ctx) error {
	db, err := utils.ConnectionDb()

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var newPerson domain.Person

	getPerson := db.Table("people").Find(&person).First(&newPerson)

	if getPerson.Error != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "person not found"})
	}

	result := db.Table("people").Where("id = ?", newPerson.ID).Update("score", 50)

	if result.Error != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "[insert score] não foi possivel inserir score para o " + person.Name})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "score entered"})
}
