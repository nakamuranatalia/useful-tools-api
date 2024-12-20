package database

import (
	"fmt"
	"log"

	"github.com/nakamuranatalia/useful-tools-api/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "postgres"
	user     = "user"
	password = "psswrd"
	dbName   = "useful-tools"
	sslMode  = "disable"
)

func DatabaseConnection() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=%s", host, user, dbName, password, sslMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panicf("Could not start the database, error: %v", err)
	}
	db.AutoMigrate(&model.Tool{})

	return db
}
