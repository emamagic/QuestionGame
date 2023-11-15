package usercontroller

import (
	"game/param"
	"game/pkg/richerror"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c Controller) userRegister(e echo.Context) error {
	var req param.RegisterRequest
	if err := e.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if fieldErrors, err := c.userValidator.ValidateUserRegister(req); err != nil {
		msg, code := richerror.Error(err)
		return e.JSON(code, echo.Map{
			"message": msg,
			"errors":  fieldErrors,
		})
	}

	resp, err := c.userSvc.Register(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusCreated, resp)
}