package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	// help the others file to interact with the DB
	db *gorm.DB
)

func Connect() {
	// username := os.Getenv("db_user")
	username := "postgres"
	password := "admin"
	dbName := "practice"
	dbHost := "localhost"
	// add port because i have 2 version postgres, and the one that i use is not on the default port
	dbPort := "5433"

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", dbHost, username, dbName, password, dbPort)

	// if there's an error, the error will come into err, if not it will come into d
	d, err := gorm.Open("postgres", dbUri)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}
