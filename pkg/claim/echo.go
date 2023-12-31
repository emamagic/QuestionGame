package claim

import (
	"game/config"
	"game/service/authservice"
	"github.com/labstack/echo/v4"
)

func GetClaimsFromEchoContext(c echo.Context) *authservice.Claims {
	// let it crash
	return c.Get(config.AuthMiddlewareContextKey).(*authservice.Claims)
}