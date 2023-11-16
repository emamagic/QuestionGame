package userservice

import (
	"game/param"
	"game/pkg/richerror"
)

func (s Service) Profile(req param.ProfileRequest) (param.ProfileResponse, error) {
	op := "userservice.Profile" 
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		// I do not expect the repository call return "record not found",
		// because I assume the interactor input is sanitized
		return param.ProfileResponse{},
			richerror.New(op).
				WithCode(richerror.CodeUnexpected).
				WithErr(err).
				WithMeta(map[string]interface{}{"req": req})

	}

	return param.ProfileResponse{Name: user.Name}, nil
}
