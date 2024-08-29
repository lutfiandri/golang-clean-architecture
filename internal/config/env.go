package config

import "github.com/spf13/viper"

var (
	APP_NAME               string
	APP_PREFORK            bool
	APP_PORT               int
	CORS_ALLOW_ORIGIN      string
	LOG_LEVEL              int
	JWT_SECRET_KEY         string
	DATABASE_USERNAME      string
	DATABASE_PASSWORD      string
	DATABASE_HOST          string
	DATABASE_PORT          int
	DATABASE_NAME          string
	DATABASE_POOL_IDLE     int
	DATABASE_POOL_MAX      int
	DATABASE_POOL_LIFETIME int
)

func LoadEnv(viper *viper.Viper) {
	APP_NAME = viper.GetString("APP_NAME")
	APP_PREFORK = viper.GetBool("APP_PREFORK")
	APP_PORT = viper.GetInt("APP_PORT")
	CORS_ALLOW_ORIGIN = viper.GetString("CORS_ALLOW_ORIGIN")
	LOG_LEVEL = viper.GetInt("LOG_LEVEL")
	JWT_SECRET_KEY = viper.GetString("JWT_SECRET_KEY")
	DATABASE_USERNAME = viper.GetString("DATABASE_USERNAME")
	DATABASE_PASSWORD = viper.GetString("DATABASE_PASSWORD")
	DATABASE_HOST = viper.GetString("DATABASE_HOST")
	DATABASE_PORT = viper.GetInt("DATABASE_PORT")
	DATABASE_NAME = viper.GetString("DATABASE_NAME")
	DATABASE_POOL_IDLE = viper.GetInt("DATABASE_POOL_IDLE")
	DATABASE_POOL_MAX = viper.GetInt("DATABASE_POOL_MAX")
	DATABASE_POOL_LIFETIME = viper.GetInt("DATABASE_POOL_LIFETIME")
}
