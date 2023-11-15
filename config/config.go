package config

import (
	"game/repository/mysql"
	"game/service/authservice"
)

type HTTPServer struct {
	Port int
}
type Config struct {
	Auth       authservice.Config
	HTTPServer HTTPServer
	Mysql      mysql.Config
}
