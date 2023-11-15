package richerror

const (
	DBError                       = "can not scan query result"
	ParsingError                  = "can not parse error"
	InvalidPhoneNumberOrPass      = "phone_number or password is not correct"
	InvalidPhoneNumber            = "phone number is not valid"
	RepetitivePhonNumber          = "phone number is not unique"
	NameLessThanThreeNumber       = "name length should be greater then 3"
	PhonNumberLessThanThreeNumber = "pass length should be greater then 3"
	RecordNotFound                = "record not found"
	InvalidInput                  = "you input is not valid"
)
