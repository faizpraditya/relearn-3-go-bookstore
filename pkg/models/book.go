package models

import (
	"go-bookstore/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// helper to intialize with the db in the first place
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
