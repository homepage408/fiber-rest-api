package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	cf "fiber-rest-api/configuration"
)

var (
	conf = cf.Configuration{}
	d    struct{}
)

func runApplication() error {
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, skipping...")
	}

	conf.Init()

	server := fiber.New(
		fiber.Config{
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"Status":    fiber.StatusInternalServerError,
					"IsSuccess": false,
					"Message":   err.Error(),
					"Data":      d,
				})
			},
			AppName: conf.Service.Name,
		},
	)

	Handler(server)
	port := "8080"
	if conf.Service.Port != 0 {
		port = strconv.Itoa(conf.Service.Port)
	}

	return server.Listen(":" + port)
}

func Handler(route *fiber.App) {

	// db := cf.NewDb(&conf)
	route.Use(logger.New(logger.Config{
		Format: "[${time}] ${locals:requestid} ${status} - ${latency} ${method} ${path}\n",
	}))
	route.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType",
	}))

	route.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "Hello World!",
		})
	})
}

func main() {
	if err := runApplication(); err != nil {
		log.Fatal(err)
	}
}
