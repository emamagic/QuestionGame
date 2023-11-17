package backofficeuserservice

import "game/domain"

type Service struct{}

func New() Service {
	return Service{}
}

func (s Service) ListAllUsers() ([]domain.User, error) {
	// TODO - implement me
	list := make([]domain.User, 0)

	list = append(list, domain.User{
		ID:           0,
		PhoneNumber:  "fake",
		Name:         "fake",
		HashPassword: "fake",
		Role:         domain.AdminRole,
	})

	return list, nil
}
