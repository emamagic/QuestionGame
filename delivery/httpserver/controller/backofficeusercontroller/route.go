package backofficeusercontroller

import (
	"game/delivery/httpserver/middleware"
	"game/domain"
	
	"github.com/labstack/echo/v4"
)

func (c Controller) SetRoutes(e *echo.Echo) {
	userGroup := e.Group("/backoffice/users")

	userGroup.GET("/", c.listUsers, middleware.Auth(c.authSvc, c.authConfig),
		middleware.AccessCheck(c.authorizationSvc, domain.UserListPermission))
}
