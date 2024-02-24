package models

import (
	"fmt"
	"go-bookstore/config"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}


func init(){
	config.Connnect()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
	db.Create(&b)
	return b
}

func GetAllBook() []*Book{
	var Books []*Book
	db.Find(Books)
	return Books
}


func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book
	db:=db.Where(fmt.Sprintf("ID=%s", Id)).Find(&getBook)
	return &getBook, db
}


func DeleteBook(ID int64) Book{
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}