package param

import "game/domain"

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginResponse struct {
	UserInfo UserInfo `json:"user"`
	Tokens   domain.Tokens   `json:"tokens"`
}
