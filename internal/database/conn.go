package database

import (
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

// func CreateDatabase() {
// 	dsn := "host=localhost user=user password=psswrd port=5432 sslmode=disable"
// 	count := 0

// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Panic("Não rolou conectar não :/")
// 	}

// 	db.Raw("SELECT count(*) FROM pg_database WHERE datname = ?", DBNAME).Scan(&count)
// 	if count == 0 {
// 		db.Exec("CREATE DATABASE %s", DBNAME)
// 	}
// }

func DatabaseConnection() *gorm.DB {
	dsn := "username=user password=psswrd dbname=useful-tools host=localhost port=5432 sslmode=disable"
	//CreateDatabase()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Error connecting to the database")
	}

	return db
}
