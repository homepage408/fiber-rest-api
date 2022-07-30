package service

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type ExampleStruct struct {
	ID     int16  `json:"id"`
	Name   string `json:"name"`
	Alamat string `json:"alamat"`
}

func PostLocal(c *fiber.Ctx) error {
	var req ExampleStruct

	if err := c.BodyParser(&req); err != nil {
		log.Fatal(err)
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": req,
	})
}
