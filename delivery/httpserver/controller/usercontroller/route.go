package usercontroller

import (
	"game/delivery/httpserver/middleware"

	"github.com/labstack/echo/v4"
)

func (h Controller) SetUserRoutes(e *echo.Echo) {
	userGroup := e.Group("/users")

	userGroup.POST("/register", h.userRegister)
	userGroup.POST("/login", h.userLogin)
	userGroup.GET("/profile", h.userProfile, middleware.Auth(h.authSvc, h.authConfig))
}
