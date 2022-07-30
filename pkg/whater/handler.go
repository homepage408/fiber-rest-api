package wheater

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type wheaterHandler struct {
	wheaterService WheaterService
}

func NewHandler(wheaterService WheaterService) *wheaterHandler {
	return &wheaterHandler{wheaterService}
}

func (wh *wheaterHandler) GetWheater(c *fiber.Ctx) error {
	inputReq := new(InputReq)

	if err := c.BodyParser(inputReq); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})
	}

	var api string
	if inputReq.Forecast {
		api = "v1/forecast.json"
	} else {
		api = "v1/current.json"
	}

	fmt.Println(api)
	fmt.Println(inputReq)
	response, err := wh.wheaterService.GetWheater(api, inputReq.Method, inputReq.City, inputReq.Forecast)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
	return c.Status(200).JSON(response)
}
