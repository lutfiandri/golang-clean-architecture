package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(viper *viper.Viper, log *zap.Logger) *gorm.DB {
	username := viper.GetString("DATABASE_USERNAME")
	password := viper.GetString("DATABASE_PASSWORD")
	host := viper.GetString("DATABASE_HOST")
	port := viper.GetInt("DATABASE_PORT")
	database := viper.GetString("DATABASE_NAME")
	idleConnection := viper.GetInt("DATABASE_POOL_IDLE")
	maxConnection := viper.GetInt("DATABASE_POOL_MAX")
	maxLifeTimeConnection := viper.GetInt("DATABASE_POOL_LIFETIME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", host, username, password, database, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: %v", zap.Error(err))
	}

	connection, err := db.DB()
	if err != nil {
		log.Fatal("failed to connect database: %v", zap.Error(err))
	}

	connection.SetMaxIdleConns(idleConnection)
	connection.SetMaxOpenConns(maxConnection)
	connection.SetConnMaxLifetime(time.Second * time.Duration(maxLifeTimeConnection))

	return db
}
