package config

import "time"

type Settings struct {
	ApiEnv          string `mapstructure:"API_ENV"`
	JWTSecretKey    string
	AccessTokenDur  time.Duration
	RefreshTokenDur time.Duration
	LogLevel        string
	LogOutput       string
}
