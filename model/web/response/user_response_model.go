package response

type WebUserResponse struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at,omitempty"`
	Token     string `json:"token,omitempty"`
}
