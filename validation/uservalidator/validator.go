package uservalidator


type Validator struct {
	svc Service
}

type Service interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
}

func New(svc Service) Validator {
	return Validator{svc: svc}
}
