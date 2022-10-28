package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtMiddleware "github.com/gofiber/jwt/v2"
)

func JWTProtected() func(*fiber.Ctx) error {
	// Create config for JWT authentication middleware.
	config := jwtMiddleware.Config{
		SigningKey:    []byte(os.Getenv("JWT_SECRET_KEY")),
		ContextKey:    "jwt", // used in private routes
		ErrorHandler:  jwtError,
		SigningMethod: "HS512",
	}

	return jwtMiddleware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	// Return status 401 and failed authentication error.
	if err.Error() == "Missing or malformed JWT" {
		return fiber.NewError(fiber.StatusBadRequest, "token is missing")
	} else if err.Error() == "Token is expired" {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized, check expiration time of your token")
	}
	return fiber.NewError(fiber.StatusUnauthorized, err.Error())
}
