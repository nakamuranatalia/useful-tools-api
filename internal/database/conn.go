package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dsn     = "host=localhost user=user password=psswrd dbname=useful-tools port=5432 sslmode=disable"
	HOST    = "localhost"
	USER    = "user"
	PSSWRD  = "psswrd"
	DBNAME  = "useful-tools"
	PORT    = "5432"
	SSLMODE = "disabled"
)

func DatabaseConnection() *gorm.DB {
	createDatabase()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Error connecting to the database")
	}

	return db
}

func createDatabase() {
	connDsn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmmode=%s", HOST, USER, PSSWRD, PORT, SSLMODE)
	count := 0

	db, err := gorm.Open(postgres.Open(connDsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Raw("SELECT count(*) FROM %s", DBNAME).Scan(&count)
	if count == 0 {
		db.Exec("CREATE DATABASE %s", DBNAME)
	}
}
