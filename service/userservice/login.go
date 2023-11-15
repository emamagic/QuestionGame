package userservice

import (
	"game/domain"
	"game/param"
	"game/pkg/richerror"
)

func (s Service) Login(req param.LoginRequest) (param.LoginResponse, error) {
	op := "userservice.Login"
	user, gErr := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if gErr != nil {
		return param.LoginResponse{}, richerror.New(op).WithErr(gErr).
			WithMeta(map[string]interface{}{"phone_number": req.PhoneNumber})
	}

	accessToken, aErr := s.auth.CreateAccessToken(user)
	if aErr != nil {
		return param.LoginResponse{},
			richerror.New(op).
				WithCode(richerror.CodeUnexpected).
				WithErr(aErr)
	}

	refreshToken, rErr := s.auth.CreateRefreshToken(user)
	if rErr != nil {
		return param.LoginResponse{},
			richerror.New(op).
				WithCode(richerror.CodeUnexpected).
				WithErr(rErr)
	}

	return param.LoginResponse{UserInfo: param.UserInfo{
		ID:          user.ID,
		PhoneNumber: user.PhoneNumber,
		Name:        user.Name,
	},
		Tokens: domain.Tokens{AccessToken: accessToken, RefreshToken: refreshToken},
	}, nil
}
