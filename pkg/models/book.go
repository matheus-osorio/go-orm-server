package models

import (
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func (book *Book) CreateBook() *Book {
	db.Create(&book)
	return book
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id int64) (Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", id).Find(&book)
	return book, db
}

func DeleteBookById(id int64) Book {
	var book Book

	db.Where("ID=?", id).Delete(&book)

	return book
}
