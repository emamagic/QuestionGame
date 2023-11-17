package userservice

import (
	"game/domain"
	"game/param"
	"game/pkg/richerror"
)

func (s Service) Register(p param.RegisterRequest) (param.RegisterResponse, error) {
	op := "userservice.Register"
	// TODO - tech debt => we should verify phone number by verification code
	hashPassword, hashPasErr := s.hashPassGen.GenerateFromPassword(p.Password)
	if hashPasErr != nil {
		return param.RegisterResponse{},
			richerror.New(op).
				WithCode(richerror.CodeUnexpected).
				WithErr(hashPasErr)
	}

	user := domain.User{
		PhoneNumber:  p.PhoneNumber,
		Name:         p.Name,
		HashPassword: hashPassword,
		Role:         domain.UserRole,
	}

	createdUser, err := s.repo.Register(user)
	if err != nil {
		return param.RegisterResponse{},
			richerror.New(op).
				WithCode(richerror.CodeUnexpected).
				WithErr(err)
	}

	return param.RegisterResponse{User: param.UserInfo{
		ID:          createdUser.ID,
		PhoneNumber: createdUser.Name,
		Name:        createdUser.PhoneNumber,
	}}, nil
}
