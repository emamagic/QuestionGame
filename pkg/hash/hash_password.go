package hash

import (
	"game/pkg/richerror"
	"golang.org/x/crypto/bcrypt"
)

type hashPassword struct{}

func New() hashPassword {
	return hashPassword{}
}

func (h hashPassword) GenerateFromPassword(password string) (string, error) {
	op := "hash.GenerateFromPassword"

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "",
			richerror.New(op).
				WithCode(richerror.CodeUnexpected).
				WithErr(err)
	}
	return string(hashPassword), nil
}

func (h hashPassword) CompareHashAndPassword(hashPassword string, password string) error {
	op := "hash.CompareHashAndPassword"

	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return richerror.New(op).
			WithCode(richerror.CodeInvalid).
			WithMessage(richerror.InvalidPhoneNumberOrPass)
	}
	return nil
}
