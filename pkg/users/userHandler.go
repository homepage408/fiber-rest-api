package users

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber"
)

type userHandler struct {
	userService UserService
}

func NewHandlerUSer(userService UserService) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) FindAllUser(c *fiber.Ctx) error {

	var userResponse []UserResponse
	user, err := h.userService.FindAllUser()
	if err != nil {
		log.Fatal(err)
		return c.Status(500).JSON(ResponseErr(err))
	}

	for _, data := range user {
		user := UserResponse{
			ID:        data.ID,
			FirstName: data.FirstName,
			LastName:  data.LastName,
			Username:  data.FirstName,
			Email:     data.Email,
			Role:      data.Role,
		}

		userResponse = append(userResponse, user)
	}

	return c.Status(200).JSON(userResponse)
}

func (h *userHandler) FindUserById(c *fiber.Ctx) error {
	userID := c.Params("id")

	id, _ := strconv.Atoi(userID)
	user, err := h.userService.FindUserById(id)
	if err != nil {
		log.Fatal(err)
		return c.Status(500).JSON(ResponseErr(err))
	}

	return c.Status(200).JSON(FuncUserResponse(user))
}

func (h *userHandler) SignUp(c *fiber.Ctx) error {
	var userIn UserRequest

	if err := c.BodyParser(&userIn); err != nil {
		log.Fatal(err)
		return c.Status(500).JSON(ResponseErr(err))
	}

	errors := ValidateStruct(userIn)
	if errors != nil {
		return c.Status(400).JSON(errors)
	}

	user, err := h.userService.SignUp(userIn)
	if err != nil {
		log.Fatal(err)
		return c.Status(500).JSON(ResponseErr(err))
	}

	return c.Status(200).JSON(FuncUserResponse(user))
}
