package userservice

import (
	"game/domain"
	"game/pkg/hash"
)

type Service struct {
	auth domain.AuthGenerator
	hashPassGen hash.HashPassGen
	repo domain.UserRepo
}

func New(authGenerator domain.AuthGenerator, repo domain.UserRepo, hashPassGen hash.HashPassGen) Service {
	return Service{auth: authGenerator, repo: repo, hashPassGen: hashPassGen}
}