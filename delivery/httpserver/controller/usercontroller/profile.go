package usercontroller

import (
	"game/param"
	"game/pkg/richerror"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c Controller) userProfile(e echo.Context) error {
	var req param.ProfileRequest
	if err := e.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	authToken := e.Request().Header.Get("Authorization")
	
	if fieldErrors, err := c.userValidator.ValidateUserProfile(authToken); err != nil {
		msg, code := richerror.Error(err)
		return e.JSON(code, echo.Map{
			"message": msg,
			"errors":  fieldErrors,
		})
	}
	// error has been check in validator
	claims, _ := c.authSvc.ParseToken(authToken)

	resp, pErr := c.userSvc.Profile(param.ProfileRequest{UserID: claims.UserID})
	if pErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, pErr.Error())
	}

	return e.JSON(http.StatusOK, resp)

}
