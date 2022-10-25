package handler

import (
	"fiber-rest-api/model/web/request"
	"fiber-rest-api/model/web/response"
	usercase "fiber-rest-api/pkg/usecase"
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
	if err := c.BodyParser(req); err != nil {
		responseApp := response.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
		}
		return c.Status(http.StatusInternalServerError).JSON(responseApp)
	}

	userResponse := handler.UserUsecase.Login(context, req)
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
	if err := c.BodyParser(req); err != nil {
		responseApp := response.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
		}
		return c.Status(http.StatusInternalServerError).JSON(responseApp)
	}

	//masuk ke usecase
	resp := handler.UserUsecase.SignUp(context, req)

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
	if err := c.BodyParser(req); err != nil {
		responseApp := response.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
		}
		return c.Status(http.StatusInternalServerError).JSON(responseApp)
	}

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
	if err := c.BodyParser(req); err != nil {
		responseApp := response.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
		}
		return c.Status(http.StatusInternalServerError).JSON(responseApp)
	}

	resp := handler.UserUsecase.FindByEmail(context, req)
	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   resp,
	}

	return c.Status(http.StatusOK).JSON(webResponse)
}
