package backofficeusercontroller

import (
	"game/pkg/richerror"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c Controller) listUsers(e echo.Context) error {
	list, err := c.backofficeUserSvc.ListAllUsers()
	if err != nil {
		msg, code := richerror.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return e.JSON(http.StatusOK, echo.Map{
		"data": list,
	})
}