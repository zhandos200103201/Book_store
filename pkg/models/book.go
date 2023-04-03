package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string
	Description string
	Author      string
	Rating      float32
	Price       float32
	Language    string
	Image       string
	Category    string
}

var db *gorm.DB

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	fmt.Println("Getting all books")
	return books
}

func GetBookById(id int16) (*Book, *gorm.DB) {
	var book Book
	db := db.Find(&book, id)
	fmt.Println("hello world")
	return &book, db
}

func DeleteBook(id int16) Book {
	var book Book
	db.Where("ID=?", id).Delete(&book)
	return book
}
