package response

import "database/sql"

type WebUserResponse struct {
	Id        int            `json:"id"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	CreatedAt sql.NullString `json:"created_at"`
	Token     string         `json:"token,omitempty"`
}
