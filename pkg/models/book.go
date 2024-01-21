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

func (b *Book) CreateBookDB() *Book {
	// from gorm
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooksDB() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookByIdDB(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBookDB(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}
