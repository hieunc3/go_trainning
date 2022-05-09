package models

import (
	"go_tranning/book_management/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json: "author"`
	Publication string `json: "publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) *Book {
	var getBook Book
	db.First(&getBook, Id)
	return &getBook
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Delete(&Book{}, Id)
	return book
}

func (b *Book) UpdateBook() *Book {
	db.Save(&b)
	return b
}
