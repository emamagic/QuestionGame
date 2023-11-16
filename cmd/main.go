package main

import (
	"game/config"
	"game/delivery/httpserver"
	"game/delivery/httpserver/controller/usercontroller"
	"game/pkg/hash"
	"game/repository/mysql"
	"game/repository/mysql/mysqluser"
	"game/service/authservice"
	"game/service/userservice"
	"game/validation/uservalidator"
	"time"
)

const (
	JwtSignKey                 = "jwt-secret"
	AccessTokenSubject         = "at"
	RefreshTokenSubject        = "rt"
	AccessTokenExpireDuration  = time.Hour * 24
	RefreshTokenExpireDuration = time.Hour * 24 * 7
)

func main() {
	cfg := config.Config{
		HTTPServer: config.HTTPServer{Port: 8080},
		Auth: authservice.Config{
			SignKey:               JwtSignKey,
			AccessExpirationTime:  AccessTokenExpireDuration,
			RefreshExpirationTime: RefreshTokenExpireDuration,
			AccessSubject:         AccessTokenSubject,
			RefreshSubject:        RefreshTokenSubject,
		},
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
	hashGen := hash.New()
	authSvc := authservice.New(cfg.Auth)
	mysql := mysql.New(cfg.Mysql)
	mysqlusers := mysqluser.New(mysql)
	userValidator := uservalidator.New(mysqlusers, authSvc, hashGen)
	userSvc := userservice.New(authSvc, mysqlusers, hashGen)
	return usercontroller.New(userSvc, authSvc, userValidator)
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
