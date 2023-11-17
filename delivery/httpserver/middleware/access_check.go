package middleware

import (
	"game/domain"
	"game/pkg/claim"
	"game/pkg/richerror"
	"game/service/authorizationservice"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AccessCheck(service authorizationservice.Service,
	permissions ...domain.PermissionTitle) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			claims := claim.GetClaimsFromEchoContext(c)
			isAllowed, err := service.CheckAccess(claims.UserID, claims.Role, permissions...)
			if err != nil {
				// TODO - log unexpected error
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"message": richerror.SomethingWentWrong,
				})
			}

			if !isAllowed {
				return c.JSON(http.StatusForbidden, echo.Map{
					"message": richerror.UserNotAllowed,
				})
			}

			return next(c)
		}
	}
}