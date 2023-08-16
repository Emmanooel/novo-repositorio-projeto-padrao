package models

import (
	"fmt"
	"log"
	"projetoPadrao/internal/domain"
	"projetoPadrao/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func CreatePerson(person *domain.Person, c *fiber.Ctx) error {
	db, err := utils.ConnectionDb()
	if err != nil {
		err := fmt.Sprintf("[Create Person] Error ao efetuar conexão com o banco de dados \n error: %v", err.Error())
		log.Print(err)
		return fiber.NewError(fiber.StatusBadRequest, err)
	}

	result := db.Table("person").Create(person)

	log.Println(result)
	return c.Status(fiber.StatusCreated).JSON(person)
}
