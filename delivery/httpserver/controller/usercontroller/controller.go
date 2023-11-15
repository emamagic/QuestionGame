package usercontroller

import (
	"game/service/userservice"
	"game/validation/uservalidator"
)

type Controller struct {
	userSvc       userservice.Service
	userValidator uservalidator.Validator
}

func New(userSvc userservice.Service, userValidator uservalidator.Validator) Controller {
	return Controller{
		userSvc:       userSvc,
		userValidator: userValidator,
	}
}
