package config

import "time"

var defaultConfig = map[string]interface{}{
	"auth.refresh_subject":                  RefreshTokenSubject,
	"auth.access_subject":                   AccessTokenSubject,
	"auth.refresh_expiration_time":          RefreshTokenExpireDuration,
	"auth.access_expiration_time":           AccessTokenExpireDuration,
	"application.graceful_shutdown_timeout": time.Second * 5,
}