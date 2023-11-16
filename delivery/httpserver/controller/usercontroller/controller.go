package usercontroller

import (
	"game/service/authservice"
	"game/service/userservice"
	"game/validation/uservalidator"
)

type Controller struct {
	userSvc       userservice.Service
	userValidator uservalidator.Validator
	authSvc       authservice.Service
}

func New(userSvc userservice.Service, authSvc authservice.Service, userValidator uservalidator.Validator) Controller {
	return Controller{
		userSvc:       userSvc,
		authSvc:       authSvc,
		userValidator: userValidator,
	}
}
