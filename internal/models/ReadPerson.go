package models

import (
	"fmt"
	"log"
	"projetoPadrao/internal/domain"
	"projetoPadrao/internal/utils"

	"github.com/gofiber/fiber/v2"
)

//get person or find person
func ReadPerson(person *domain.Person, c *fiber.Ctx) error {
	db, err := utils.ConnectionDb()

	if err != nil {
		err := fmt.Sprintf("[Read Person] Error ao efetuar conexão com o banco de dados \n error: %v", err.Error())
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, err)
	}
	var user domain.Person
	result := db.Table("person").Where(&person).First(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "[Read person] pessoa não encontrada"})
	}

	log.Println(user)
	return c.Status(fiber.StatusCreated).JSON(user)
}

//get person by login
func FindByLogin(login *domain.Login, c *fiber.Ctx) error {
	db, err := utils.ConnectionDb()

	if err != nil {
		err := fmt.Sprintf("[Find by Email] Error ao efetuar conexão com o banco de dados \n error: %v", err.Error())
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, err)
	}

	var user domain.Person
	result := db.Table("person").Where(&login).First(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "[Find by Email] Login não encontrado"})
	}
	log.Println(user)
	return c.Status(fiber.StatusOK).JSON(user)
}

//get ranking
func Ranking(c *fiber.Ctx) error {
	db, err := utils.ConnectionDb()

	if err != nil {
		fm := fmt.Sprintf("[ranking] Error ao conectar com o banco de dados \n error: %v", err.Error())
		return fiber.NewError(fiber.StatusBadRequest, fm)
	}

	var ranking []domain.Ranking
	result := db.Table("person").
		Select("id, name, score").
		Order("score DESC").Scan(&ranking)

	if result.Error != nil {
		return fiber.NewError(fiber.StatusBadRequest, "[Ranking] error ao consultar banco de dados")
	}

	return c.Status(fiber.StatusOK).JSON(ranking)
}
