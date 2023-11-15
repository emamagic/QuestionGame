package userservice

import (
	"game/domain"
	"game/pkg/hash"
)

type Repository interface {
	Register(u domain.User) (domain.User, error)
	GetUserByPhoneNumber(phoneNumber string) (domain.User, error)
}

type Service struct {
	auth domain.AuthGenerator
	hashPassGen hash.HashPassGen
	repo Repository
}

func New(authGenerator domain.AuthGenerator, repo Repository, hashPassGen hash.HashPassGen) Service {
	return Service{auth: authGenerator, repo: repo, hashPassGen: hashPassGen}
}