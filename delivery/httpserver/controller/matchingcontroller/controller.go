package matchingcontroller

import (
	"game/service/authservice"
	"game/service/matchingservice"
	"game/validation/matchingvalidator"
)

type Controller struct {
	authConfig        authservice.Config
	authSvc           authservice.Service
	matchingSvc       matchingservice.Service
	matchingValidator matchingvalidator.Validator
}

func New(
	authConfig authservice.Config,
	authSvc authservice.Service,
	matchingSvc matchingservice.Service,
	matchingValidator matchingvalidator.Validator,
) Controller {
	return Controller{
		authConfig:        authConfig,
		authSvc:           authSvc,
		matchingSvc:       matchingSvc,
		matchingValidator: matchingValidator,
	}
}
