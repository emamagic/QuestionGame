package matchingvalidator

import (
	"fmt"
	"game/domain"
	"game/param"
	"game/pkg/richerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateAddToWaitingListRequest(req param.AddToWaitingListRequest) (map[string]string, error) {
	const op = "matchingvalidator.AddToWaitingListRequest"

	if err := validation.ValidateStruct(&req,

		validation.Field(&req.Category,
			validation.Required,
			validation.By(v.isCategoryValid)),
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

func (v Validator) isCategoryValid(value interface{}) error {
	category := value.(domain.Category)

	if !category.IsValid() {
		return fmt.Errorf(richerror.CategoryIsNotValid)
	}

	return nil
}
