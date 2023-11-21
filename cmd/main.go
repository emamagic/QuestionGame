package main

import (
	"game/adapter/redis"
	"game/config"
	"game/delivery/httpserver"
	"game/delivery/httpserver/controller/backofficeusercontroller"
	"game/delivery/httpserver/controller/matchingcontroller"
	"game/delivery/httpserver/controller/usercontroller"
	"game/pkg/hash"
	"game/repository/migrator"
	"game/repository/mysql"
	"game/repository/mysql/mysqlaccesscontrol"
	"game/repository/mysql/mysqluser"
	"game/repository/redis/redismatching"
	"game/service/authorizationservice"
	"game/service/authservice"
	"game/service/backofficeuserservice"
	"game/service/matchingservice"
	"game/service/userservice"
	"game/validation/matchingvalidator"
	"game/validation/uservalidator"
)

func main() {

	cfg := config.Load("config.yml")

	// TODO - add command for migration
	mgr := migrator.New(cfg.Mysql)
	mgr.Up()

	userController, backofficeusercontroller, matchingcontroller := setupServices(cfg)
	server := httpserver.New(cfg, userController, backofficeusercontroller, matchingcontroller)

	server.Serve()

}

func setupServices(cfg config.Config) (usercontroller.Controller, backofficeusercontroller.Controller, matchingcontroller.Controller) {
	hashGen := hash.New()

	authSvc := authservice.New(cfg.Auth)

	mysql := mysql.New(cfg.Mysql)
	mysqlusers := mysqluser.New(mysql)
	mysqlaccesscontrol := mysqlaccesscontrol.New(mysql)

	userValidator := uservalidator.New(mysqlusers, authSvc, hashGen)

	userSvc := userservice.New(authSvc, mysqlusers, hashGen)
	usercontroller := usercontroller.New(userSvc, authSvc, userValidator, cfg.Auth)

	authorizationSvc := authorizationservice.New(mysqlaccesscontrol)

	backofficeuserSvc := backofficeuserservice.New()
	backofficeusercontroller := backofficeusercontroller.New(
		cfg.Auth,
		authSvc,
		backofficeuserSvc,
		authorizationSvc,
	)

	redisAdapter := redis.New(cfg.Redis)
	matchingRepo := redismatching.New(redisAdapter)
	matchingvalidator := matchingvalidator.New()
	matchingservice := matchingservice.New(cfg.MatchingService, matchingRepo)
	matchingcontroller := matchingcontroller.New(cfg.Auth, authSvc, matchingservice, matchingvalidator)

	return usercontroller, backofficeusercontroller, matchingcontroller
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
