package uservalidator

import (
	"game/pkg/richerror"
)

func (v Validator) ValidateUserProfile(authToken string) (map[string]string, error) {
	const op = "uservalidator.ValidateUserProfile"

	_, paErr := v.tokenValidator.ParseToken(authToken)
	if paErr != nil {
		fieldErrors := make(map[string]string)
		fieldErrors["userID"] = "unuthorized"
		richerr := richerror.New(op).WithMessage(richerror.InvalidInput).
			WithCode(richerror.CodeInvalid).
			WithErr(paErr)

		return fieldErrors, richerr
	}

	return nil, nil

}
