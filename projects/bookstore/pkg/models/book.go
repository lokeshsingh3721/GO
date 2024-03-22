package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lokeshsingh3721/bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name string 
	Author string 
	Publication string 
}

func init (){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book
	db:=db.Where("ID=?",Id).Find(&getBook)
	return &getBook , db
}

func DeleteBook(ID int64) Book{
	var book Book
	db.Where("ID=?",ID).Delete(book)
	return book
}

