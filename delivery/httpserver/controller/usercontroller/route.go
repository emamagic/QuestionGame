package usercontroller

import "github.com/labstack/echo/v4"


func (h Controller) SetUserRoutes(e *echo.Echo) {
	userGroup := e.Group("/users")
	
	userGroup.POST("/register", h.userRegister)
}