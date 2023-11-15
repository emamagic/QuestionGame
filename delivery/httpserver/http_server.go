package httpserver

import (
	"fmt"
	"game/config"
	"game/delivery/httpserver/controller/usercontroller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config      config.Config
	userController usercontroller.Controller
}

func New(config config.Config, userController usercontroller.Controller) Server {
	return Server{config: config, userController: userController}
}

func (s Server) Serve() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health-check", s.healthCheck)

	s.userController.SetUserRoutes(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HTTPServer.Port)))
}
