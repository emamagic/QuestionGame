package uservalidator

import (
	"game/domain"
	"game/pkg/hash"
	"game/service/authservice"
)

const (
	phoneNumberRegex = "^09[0-9]{9}$"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	GetUserByPhoneNumber(phoneNumber string) (domain.User, error)
	GetUserByID(userID uint) (domain.User, error)
}

type TokenValidator interface {
	ParseToken(bearerToken string) (*authservice.Claims, error)
}

type Validator struct {
	repo            Repository
	tokenValidator  TokenValidator
	hashPassCompare hash.HashPassCompare
}

func New(repo Repository, tokenValidator TokenValidator, hashPassCompare hash.HashPassCompare) Validator {
	return Validator{repo: repo, tokenValidator: tokenValidator, hashPassCompare: hashPassCompare}
}
