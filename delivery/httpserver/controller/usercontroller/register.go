package usercontroller

import (
	"game/param"
	"game/pkg/richerror"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Controller) userRegister(c echo.Context) error {
	var req param.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if fieldErrors, err := h.userValidator.ValidateUserRegister(req); err != nil {
		msg, code := richerror.Error(err)
		return c.JSON(code, echo.Map{
			"message": msg,
			"errors":  fieldErrors,
		})
	}

	resp, err := h.userSvc.Register(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, resp)
}