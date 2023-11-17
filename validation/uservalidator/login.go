package uservalidator

import (
	"game/param"
	"game/pkg/richerror"
	"regexp"

	"github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateUserLogin(req param.LoginRequest) (map[string]string, error) {
	const op = "uservalidator.ValidateUserLogin"
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.PhoneNumber,
			validation.Required,
			validation.Match(regexp.MustCompile(phoneNumberRegex)).Error(richerror.InvalidPhoneNumberOrPass),
			validation.By(v.doesPhoneNumberExist)),

		validation.Field(&req.Password, validation.Required,
			validation.When(v.compareHashAndPassword(req.PhoneNumber, req.Password), validation.Required).Else(validation.Nil.Error(richerror.InvalidInput))),
	); err != nil {
		fieldErrors := make(map[string]string)

		errV, ok := err.(validation.Errors)
		if ok {
			for key, value := range errV {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}

		return fieldErrors, richerror.New(op).WithMessage(richerror.InvalidInput).
			WithCode(richerror.CodeInvalid).
			WithMeta(map[string]interface{}{"req": req}).WithErr(err)
	}
	
	return nil, nil
}

func (v Validator) doesPhoneNumberExist(value interface{}) error {
	phoneNumber := value.(string)
	_, err := v.repo.GetUserByPhoneNumber(phoneNumber)
	if err != nil {
		return err
	}

	return nil
}

func (v Validator) compareHashAndPassword(phoneNumber, password string) bool {
	user, uErr := v.repo.GetUserByPhoneNumber(phoneNumber)
	if uErr != nil {
		return false
	}
	err := v.hashPassCompare.CompareHashAndPassword(user.HashPassword, password)
	return err == nil
}
