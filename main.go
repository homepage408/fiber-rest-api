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
	"fiber-rest-api/middleware"
	"fiber-rest-api/pkg/handler"
	"fiber-rest-api/pkg/repository"
	"fiber-rest-api/pkg/usecase"
)

var (
	conf = cf.Configuration{}
)

func runApplication() error {
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, skipping...")
	}

	conf.Init()

	server := fiber.New(
		fiber.Config{
			ErrorHandler: exception.DefaultErrorHandler,
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

	route.Use(recover.New())
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

	// from environment
	SECRET_KEY := os.Getenv("JWT_SECRET_KEY")
	EXPIRESJWT := os.Getenv("JWT_EXPIRES")

	//Repository
	userRepository := repository.NewUserRepository()

	// Usecase
	userUsecase := usecase.NewUserUsecase(userRepository, db, SECRET_KEY, EXPIRESJWT)

	// Handler
	userHandler := handler.NewUserHandler(userUsecase)

	api := route.Group("/api")

	api.Post("/login", userHandler.Login)
	api.Post("/Signup", userHandler.SignUp)
	api.Get("/FindByEmail", middleware.JWTProtected(), userHandler.FindByEmail)

}

// func Auth(c *fiber.Ctx) error {
// 	user := c.Locals("user").(*jwt.Token)
// 	claims := user.Claims.(jwt.MapClaims)
// 	name := claims["name"].(string)
// 	return c.Status(200).JSON(fiber.Map{
// 		"name": name,
// 	})
// }

// func customKeyFunc() jwt.Keyfunc {
// 	return func(t *jwt.Token) (interface{}, error) {
// 		// Always check the signing method
// 		if t.Method.Alg() != jwtware.HS256 {
// 			// fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
// 			return nil, fiber.NewError(400, "Invalid")
// 		}

// 		// TODO custom implementation of loading signing key like from a database
// 		signingKey := "6hbqvbH8UWQkHwMzxV2QCq8GyXRJ8NGREZprJeKXek3FcG"

// 		return []byte(signingKey), nil
// 	}
// }

func main() {
	if err := runApplication(); err != nil {
		log.Fatal(err)
	}
}
