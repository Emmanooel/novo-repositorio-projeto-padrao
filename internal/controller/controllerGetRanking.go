package controller

import (
	"projetoPadrao/internal/models"

	"github.com/gofiber/fiber/v2"
)

func ControllerGetRanking(c *fiber.Ctx) error {
	ranking := models.Ranking(c)

	return ranking
}
