package uservalidator

import (
	"game/domain"
	"game/pkg/hash"
)

const (
	phoneNumberRegex = "^09[0-9]{9}$"
)

type Service interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	GetUserByPhoneNumber(phoneNumber string) (domain.User, error)
}

type Validator struct {
	svc Service
	hashPassCompare hash.HashPassCompare
}

func New(svc Service, hashPassCompare hash.HashPassCompare) Validator {
	return Validator{svc: svc, hashPassCompare: hashPassCompare}
}
