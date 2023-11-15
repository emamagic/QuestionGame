package richerror

import (
	"net/http"
)

func Error(err error) (message string, code int) {

	switch er := err.(type) {
	case ErrorResponse:
		msg := er.Message()

		code := mapKindToHTTPStatusCode(er.Code())

		// we should not expose unexpected error messages
		if code >= 500 {
			msg = "something went wrong"
		}

		return msg, code
	default:
		return err.Error(), http.StatusBadRequest
	}
}

func mapKindToHTTPStatusCode(code Code) int {
	switch code {
	case CodeInvalid:
		return http.StatusUnprocessableEntity
	case CodeNotFound:
		return http.StatusNotFound
	case CodeForbidden:
		return http.StatusForbidden
	case CodeUnexpected:
		return http.StatusInternalServerError
	default:
		return http.StatusBadRequest
	}
}
