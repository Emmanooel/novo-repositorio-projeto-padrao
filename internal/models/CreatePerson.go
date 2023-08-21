package models

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"projetoPadrao/internal/domain"
	"projetoPadrao/internal/utils"
)

func CreatePerson(person *domain.Person, c *fiber.Ctx) error {
	db, err := utils.ConnectionDb()
	if err != nil {
		err := fmt.Sprintf("[Create Person] Error ao conectar com o banco de dados \n error: %v", err.Error())
		log.Print(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	result := db.Table("person").Create(person)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(&person)
}
