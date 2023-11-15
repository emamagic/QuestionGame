package richerror

import (
	"encoding/json"
)

// https://github.com/microsoft/api-guidelines/blob/vNext/Guidelines.md
type Code int

const (
	CodeInvalid Code = iota + 1
	CodeForbidden
	CodeNotFound
	CodeUnexpected
)

func (e ErrorResponse) Error() string {
	return e.Message()
}

type ErrorResponse struct {
	operation  string
	code       Code
	message    string
	wrappedErr error
	meta       map[string]interface{}
}

func New(op string) ErrorResponse {
	return ErrorResponse{operation: op}
}

func (e ErrorResponse) WithCode(code Code) ErrorResponse {
	e.code = code
	return e
}

func (e ErrorResponse) WithMeta(meta map[string]interface{}) ErrorResponse {
	e.meta = meta
	return e
}

func (e ErrorResponse) WithOp(operation string) ErrorResponse {
	e.operation = operation
	return e
}

func (e ErrorResponse) WithErr(err error) ErrorResponse {
	_, ok := err.(ErrorResponse)
	if !ok {
		e.wrappedErr = New(e.operation).
			WithMessage(err.Error()).
			WithCode(e.code)
		return e
	}
	e.wrappedErr = err
	return e
}

func (e ErrorResponse) WithMessage(message string) ErrorResponse {
	e.message = message
	return e
}

func (e ErrorResponse) Code() Code {
	if e.code != 0 {
		return e.code
	}

	er, ok := e.wrappedErr.(ErrorResponse)
	if !ok {
		return 0
	}

	return er.Code()

}

func (e ErrorResponse) Message() string {
	if e.message != "" {
		return e.message
	}

	er, ok := e.wrappedErr.(ErrorResponse)
	if !ok {
		return e.wrappedErr.Error()
	}

	return er.Message()
}

func (e ErrorResponse) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		Operation  string
		Code       Code
		Message    string
		WrappedErr error
		Meta       map[string]interface{}
	}{
		Operation:  e.operation,
		Code:       e.code,
		Message:    e.message,
		WrappedErr: e.wrappedErr,
		Meta:       e.meta,
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}
