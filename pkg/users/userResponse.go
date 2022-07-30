package users

import "github.com/go-playground/validator/v10"

type ResponseError struct {
	message error
}

type UserResponse struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ResponseErr(err error) ResponseError {
	return ResponseError{message: err}
}

func FuncUserResponse(user User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
	}
}

var validate = validator.New()

func ValidateStruct(user UserRequest) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
