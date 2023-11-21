package matchingservice

import (
	"game/domain"
	"game/param"
	"game/pkg/richerror"
	"time"
)

type Service struct {
	config Config
	repo   Repo
}

type Repo interface {
	AddToWaitingList(userID uint, category domain.Category) error
}

type Config struct {
	WatingTimeout time.Duration `koanf:"waiting_timeout"`
}

func New(cfg Config, repo Repo) Service {
	return Service{
		config: cfg,
		repo:   repo,
	}
}

func (s Service) AddToWaitingList(req param.AddToWaitingListRequest) (param.AddToWaitingListResponse, error) {
	op := "matchingservice.AddToWaitingList"

	err := s.repo.AddToWaitingList(req.UserID, req.Category)
	if err != nil {
		return param.AddToWaitingListResponse{}, richerror.New(op).WithErr(err).WithCode(richerror.CodeUnexpected)
	}

	return param.AddToWaitingListResponse{Timeout: s.config.WatingTimeout}, nil

}
