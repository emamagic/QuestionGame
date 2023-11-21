package matchingcontroller

import (
	"game/param"
	"game/pkg/claim"
	"game/pkg/richerror"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c Controller) AddToWaitingList(e echo.Context) error {
	var req param.AddToWaitingListRequest
	if err := e.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadGateway)
	}

	claims := claim.GetClaimsFromEchoContext(e)
	req.UserID = claims.UserID

	if fieldErrors, err := c.matchingValidator.ValidateAddToWaitingListRequest(req); err != nil {
		msg, code := richerror.Error(err)
		return e.JSON(code, echo.Map{
			"message": msg,
			"errors":  fieldErrors,
		})
	}

	resp, err := c.matchingSvc.AddToWaitingList(req)
	if err != nil {
		msg, code := richerror.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return e.JSON(http.StatusOK, resp)

}
