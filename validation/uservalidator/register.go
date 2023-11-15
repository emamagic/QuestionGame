package uservalidator

import (
	"fmt"
	"game/param"
	"game/pkg/richerror"
	"regexp"

	"github.com/go-ozzo/ozzo-validation/v4"
)

const (
	phoneNumberRegex = "^09[0-9]{9}$"
)

func (v Validator) ValidateUserRegister(req param.RegisterRequest) (map[string]string, error) {
	const op = "uservalidator.ValidateRegisterRequest"

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Name,
			validation.Required,
			validation.Length(3, 50)),

		validation.Field(&req.Password,
			validation.Required,
			validation.Length(3, 50)),

		validation.Field(&req.PhoneNumber,
			validation.Required,
			validation.Match(regexp.MustCompile(phoneNumberRegex)).
				Error(richerror.InvalidInput),
			validation.By(v.checkPhoneNumberUniqueness)),
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

func (v Validator) checkPhoneNumberUniqueness(value interface{}) error {
	phoneNumber := value.(string)

	if isUnique, err := v.svc.IsPhoneNumberUnique(phoneNumber); err != nil || !isUnique {
		if err != nil {
			return err
		}

		if !isUnique {
			return fmt.Errorf(richerror.RepetitivePhonNumber)
		}
	}

	return nil
}
