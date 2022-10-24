package helper

import (
	"fiber-rest-api/model/domain"
	"fiber-rest-api/model/web/response"
)

func LoginUserResponse(user domain.User, token string) response.WebUserResponse {
	return response.WebUserResponse{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		Token:     token,
	}
}

func UserResponse(user domain.User) response.WebUserResponse {
	return response.WebUserResponse{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}
