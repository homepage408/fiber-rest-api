package handler

import (
	"fiber-rest-api/helper"
	"fiber-rest-api/model/web/request"
	"fiber-rest-api/model/web/response"
	usercase "fiber-rest-api/pkg/usecase"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	Login(c *fiber.Ctx) error
	SignUp(c *fiber.Ctx) error
	RemoveAccount(c *fiber.Ctx) error
	FindByEmail(c *fiber.Ctx) error
}

type ClientUserHandler struct {
	UserUsecase usercase.UserUsecase
}

func NewUserHandler(userUsecase usercase.UserUsecase) UserHandler {
	return &ClientUserHandler{
		UserUsecase: userUsecase,
	}
}

func (handler *ClientUserHandler) Login(c *fiber.Ctx) error {
	context := c.Context()
	req := new(request.UserLoginRequest)
	err := c.BodyParser(&req)
	helper.PanicIfError(err)

	userResponse, err := handler.UserUsecase.Login(context, req)
	if err != nil {
		return err
	}

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
	}

	return c.Status(http.StatusOK).JSON(webResponse)
}

func (handler *ClientUserHandler) SignUp(c *fiber.Ctx) error {
	context := c.Context()
	req := new(request.UserSignUpRequest)
	err := c.BodyParser(req)
	helper.PanicIfError(err)

	//masuk ke usecase
	resp, err := handler.UserUsecase.SignUp(context, req)
	log.Println(err)
	if err != nil {
		return err
	}

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   resp,
	}

	return c.Status(http.StatusOK).JSON(webResponse)
}

func (handler *ClientUserHandler) RemoveAccount(c *fiber.Ctx) error {
	context := c.Context()
	req := new(request.UserRemoveAccountRequest)
	err := c.BodyParser(req)
	helper.PanicIfError(err)

	// usecase
	handler.UserUsecase.RemoveAccount(context, req)
	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}
	return c.Status(http.StatusOK).JSON(webResponse)
}

func (handler *ClientUserHandler) FindByEmail(c *fiber.Ctx) error {
	context := c.Context()
	req := new(request.UserLoginRequest)
	err := c.BodyParser(req)
	helper.PanicIfError(err)

	resp := handler.UserUsecase.FindByEmail(context, req)
	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   resp,
	}

	return c.Status(http.StatusOK).JSON(webResponse)
}
