package httpserver

import (
	"fmt"
	"game/config"
	"game/delivery/httpserver/controller/backofficeusercontroller"
	"game/delivery/httpserver/controller/usercontroller"
	"game/service/authorizationservice"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config                   config.Config
	userController           usercontroller.Controller
	backofficeusercontroller backofficeusercontroller.Controller
	authorizationservice     authorizationservice.Service
}

func New(
	config config.Config,
	userController usercontroller.Controller,
	backofficontroller backofficeusercontroller.Controller,
	authorizationservice authorizationservice.Service,
) Server {
	return Server{
		config:                   config,
		userController:           userController,
		backofficeusercontroller: backofficontroller,
		authorizationservice:     authorizationservice,
	}
}

func (s Server) Serve() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health-check", s.healthCheck)

	s.userController.SetRoutes(e)
	s.backofficeusercontroller.SetRoutes(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HTTPServer.Port)))
}
