package backofficeusercontroller

import (
	"game/service/authorizationservice"
	"game/service/authservice"
	"game/service/backofficeuserservice"
)

type Controller struct {
	authConfig        authservice.Config
	authSvc           authservice.Service
	authorizationSvc  authorizationservice.Service
	backofficeUserSvc backofficeuserservice.Service
}

func New(
	authConfig authservice.Config,
	authSvc authservice.Service,
	backofficeUserSvc backofficeuserservice.Service,
	authorizationSvc authorizationservice.Service,
) Controller {
	return Controller{
		authConfig:        authConfig,
		authSvc:           authSvc,
		backofficeUserSvc: backofficeUserSvc,
		authorizationSvc:  authorizationSvc,
	}
}
