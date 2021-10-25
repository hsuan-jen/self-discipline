package response

import (
	model "self-discipline/model/user"
)

type UserResponse struct {
	User model.User `json:"user"`
}

type LoginResponse struct {
	User      model.User `json:"user"`
	Token     string     `json:"token"`
	ExpiresAt int64      `json:"expiresAt"`
}
