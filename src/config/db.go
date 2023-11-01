package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB
var e error

func DatabaseInit() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, dbName)
	database, e = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if e != nil {
		panic(e)
	}
}

func DB() *gorm.DB {
	return database
}
