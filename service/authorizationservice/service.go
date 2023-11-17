package authorizationservice

import (
	"game/domain"
	"game/pkg/richerror"
)

type Service struct {
	repo domain.Repository
}

func New(repo domain.Repository) Service {
	return Service{repo: repo}
}

func (s Service) CheckAccess(userID uint, role domain.Role, permissions ...domain.PermissionTitle) (bool, error) {
	const op = "authorizationservice.CheckAccess"

	permissionTitles, err := s.repo.GetUserPermissionTitles(userID, role)
	if err != nil {
		return false, richerror.New(op).WithErr(err)
	}

	for _, pt := range permissionTitles {
		for _, p := range permissions {
			if p == pt {
				return true, nil
			}
		}
	}

	return false, nil
}