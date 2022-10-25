package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"

	cf "fiber-rest-api/configuration"
	"fiber-rest-api/helper/exception"
	"fiber-rest-api/pkg/handler"
	"fiber-rest-api/pkg/repository"
	"fiber-rest-api/pkg/usecase"
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
			ErrorHandler: exception.ErrorHandler,
			AppName:      conf.Service.Name,
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

	db := cf.NewDb(&conf)

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

	route.Use(recover.New())

	// from environment
	defaultSalt := os.Getenv("SALT_DEFAULT")
	saltText := os.Getenv("SALT_TEXT")

	//Repository
	userRepository := repository.NewUserRepository()

	// Usecase
	userUsecase := usecase.NewUserUsecase(userRepository, db, defaultSalt, saltText)

	// Handler
	userHandler := handler.NewUserHandler(userUsecase)

	api := route.Group("/api")
	api.Post("/login", userHandler.Login)
	api.Post("/Signup", userHandler.SignUp)
	api.Get("/FindByEmail", userHandler.FindByEmail)

}

func main() {
	if err := runApplication(); err != nil {
		log.Fatal(err)
	}
}
