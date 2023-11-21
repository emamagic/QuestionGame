package matchingcontroller

import (
	"game/delivery/httpserver/middleware"
	
	"github.com/labstack/echo/v4"
)

func (c Controller) SetRoutes(e *echo.Echo) {
	userGroup := e.Group("/matching")

	userGroup.POST("/add-to-waiting-list", c.AddToWaitingList, middleware.Auth(c.authSvc, c.authConfig))
}
