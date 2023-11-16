package uservalidator

import (
	"game/domain"
	"game/pkg/hash"
	"game/service/authservice"
)

const (
	phoneNumberRegex = "^09[0-9]{9}$"
)

type Service interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	GetUserByPhoneNumber(phoneNumber string) (domain.User, error)
	GetUserByID(userID uint) (domain.User, error)
}

type Validator struct {
	svc             Service
	authSvc         authservice.Service
	hashPassCompare hash.HashPassCompare
}

func New(svc Service, authSvc authservice.Service, hashPassCompare hash.HashPassCompare) Validator {
	return Validator{svc: svc, authSvc: authSvc, hashPassCompare: hashPassCompare}
}
