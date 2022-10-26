package usecase

import (
	"context"
	"database/sql"
	"fiber-rest-api/helper"
	"fiber-rest-api/model/domain"
	"fiber-rest-api/model/web/request"
	"fiber-rest-api/model/web/response"
	"fiber-rest-api/pkg/repository"
	"fiber-rest-api/service"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserUsecase interface {
	Login(ctx context.Context, request *request.UserLoginRequest) response.WebUserResponse
	SignUp(ctx context.Context, request *request.UserSignUpRequest) (response.WebUserResponse, error)
	RemoveAccount(ctx context.Context, request *request.UserRemoveAccountRequest)
	FindByEmail(ctx context.Context, request *request.UserLoginRequest) response.WebUserResponse
}

type ClientUserUsecase struct {
	UserRespository repository.UserRespository
	DB              *sql.DB
	Validate        *validator.Validate
	DefaultSalt     string
	SaltText        string
}

func NewUserUsecase(userRespository repository.UserRespository, db *sql.DB, defaultSalt, saltText string) UserUsecase {
	return &ClientUserUsecase{
		UserRespository: userRespository,
		DB:              db,
		DefaultSalt:     defaultSalt,
		SaltText:        saltText,
	}
}

func (usecase *ClientUserUsecase) Login(ctx context.Context, request *request.UserLoginRequest) response.WebUserResponse {
	tx, err := usecase.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// cek password  benar atau tidak
	userInput := domain.User{
		Email: request.Email,
	}

	user, err := usecase.UserRespository.Login(ctx, tx, userInput)
	helper.PanicIfError(err)

	// mengolah token
	token := "INI TOKEN"

	return helper.LoginUserResponse(user, token)
}

func (usecase *ClientUserUsecase) SignUp(ctx context.Context, request *request.UserSignUpRequest) (response.WebUserResponse, error) {
	tx, err := usecase.DB.Begin()
	log.Println("ERR TX", err)
	helper.PanicIfError(err)
	// defer helper.CommitOrRollback(tx)

	userInput := domain.User{
		Email: request.Email,
	}

	//cek apa sudah ada email yang dipakai
	exist, _ := usecase.UserRespository.FindByEmail(ctx, tx, userInput)
	if exist != (domain.User{}) {
		return response.WebUserResponse{}, fiber.NewError(fiber.StatusBadRequest, "Email sudah digunakan")
	}

	passHash, err := service.HashPassword(request.Password, usecase.DefaultSalt, usecase.SaltText)
	helper.PanicIfError(err)

	//proses password
	userInput.Username = request.Username
	userInput.Password = passHash

	result := usecase.UserRespository.Create(ctx, tx, userInput)

	return helper.UserResponse(result), nil
}

func (usecase *ClientUserUsecase) RemoveAccount(ctx context.Context, request *request.UserRemoveAccountRequest) {
	tx, err := usecase.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userInput := domain.User{
		Email: request.Email,
	}

	usecase.UserRespository.Delete(ctx, tx, userInput)
}

func (usecase *ClientUserUsecase) FindByEmail(ctx context.Context, request *request.UserLoginRequest) response.WebUserResponse {
	tx, err := usecase.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userInput := domain.User{
		Email: request.Email,
	}

	result, err := usecase.UserRespository.FindByEmail(ctx, tx, userInput)
	helper.PanicIfError(err)

	return helper.UserResponse(result)
}
