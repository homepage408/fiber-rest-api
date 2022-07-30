package main

import (
	"encoding/json"
	"fiber-rest-api/configuration"
	"fiber-rest-api/service"
	"log"
	"os"

	wheater "fiber-rest-api/pkg/whater"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	configuration.DbInit()
	configuration.RunMigrations()

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// serviceSignature, adsad := service.Signature()
	// fmt.Println(serviceSignature)
	// fmt.Println(adsad)

	// decrypt := service.SignatureDecrypt()
	// fmt.Println(decrypt)

	// app.Listen(":3000")

	// service.GetOtp("081906889041")

	// randomUnix := time.Now().UnixNano()
	// fmt.Println(randomUnix)
	// stringRand := service.RandStringBytes(15)
	// fmt.Println(stringRand)
	// service.SignatureWithRsa()

	// service.GetWheaterService("Yogyakarta")
	// response := service.GetWheaterService("Yogyakarta")
	// fmt.Println(response)

	godotenv.Load()
	apiKey := os.Getenv("API_KEY_WHATERAPI")
	newService := wheater.NewWheaterService(apiKey)
	newhandler := wheater.NewHandler(newService)

	app.Post("/cuaca", newhandler.GetWheater)

	app.Get("/cuaca/:city", service.GetWheaterService)
	log.Fatal(app.Listen(":3000"))
}
