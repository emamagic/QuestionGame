package param

import "game/domain"

type RegisterRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type RegisterResponse struct {
	User     domain.UserInfo `json:"user"`
	Metadata string          `json:"meta_data"`
}