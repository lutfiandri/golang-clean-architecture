package infrastructure

import (
	"fmt"
	"time"

	"github.com/lutfiandri/golang-clean-architecture/internal/config"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(log *zap.Logger) *gorm.DB {
	username := config.DATABASE_USERNAME
	password := config.DATABASE_PASSWORD
	host := config.DATABASE_HOST
	port := config.DATABASE_PORT
	database := config.DATABASE_NAME
	idleConnection := config.DATABASE_POOL_IDLE
	maxConnection := config.DATABASE_POOL_MAX
	maxLifeTimeConnection := config.DATABASE_POOL_LIFETIME

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

func NewConnection(user, password, host string, port int) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateDatabase(db *gorm.DB, dbName string) error {
	return db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName)).Error
}

func DeleteDatabase(db *gorm.DB, dbName string) error {
	return db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", dbName)).Error
}
