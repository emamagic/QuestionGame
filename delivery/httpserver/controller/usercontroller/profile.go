package usercontroller

import (
	"game/param"
	claim "game/pkg/cliam"
	"game/pkg/richerror"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c Controller) userProfile(e echo.Context) error {
	claims := claim.GetClaimsFromEchoContext(e)

	resp, err := c.userSvc.Profile(
		param.ProfileRequest{UserID: claims.UserID})
	if err != nil {
		msg, code := richerror.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return e.JSON(http.StatusOK, resp)

}
