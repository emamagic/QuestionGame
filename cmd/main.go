package main

import (
	"game/config"
	"game/delivery/httpserver"
	"game/delivery/httpserver/controller/usercontroller"
	"game/pkg/hash"
	"game/repository/mysql"
	"game/repository/mysql/mysqluser"
	"game/service/userservice"
	"game/validation/uservalidator"
)

func main() {
	cfg := config.Config{
		HTTPServer: config.HTTPServer{Port: 8080},
		Mysql: mysql.Config{
			Username: "gameapp",
			Password: "gameappt0lk2o20",
			Port:     3306,
			Host:     "localhost",
			DBName:   "gameapp_db",
		},
	}

	userController := setupServices(cfg)
	server := httpserver.New(cfg, userController)

	server.Serve()

}

func setupServices(cfg config.Config) usercontroller.Controller {
	mysql := mysql.New(cfg.Mysql)
	mysqlusers := mysqluser.New(mysql)
	userValidator := uservalidator.New(mysqlusers)
	hashGen := hash.New()
	userSvc := userservice.New(mysqlusers, hashGen)
	return usercontroller.New(userSvc, userValidator)
}

// func writeTypedError(w http.ResponseWriter, code int, domainErr DomainError) {
//     errPayload := &httpError{
//         Code: domainErr.Code(),
//         Message: domainErr.Message(),
//     }

//     body, _ := json.Marshal(errPayload)
//     w.WriteHeader(code)
//     w.Write(body)
// }
