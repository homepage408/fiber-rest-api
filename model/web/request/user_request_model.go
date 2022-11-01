package request

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password"`
}

type UserSignUpRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRemoveAccountRequest struct {
	Email string `json:"email"`
}
