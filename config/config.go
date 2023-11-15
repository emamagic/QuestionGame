package config

import "game/repository/mysql"

type HTTPServer struct {
	Port int
}
type Config struct {
	HTTPServer HTTPServer
	Mysql      mysql.Config
}
