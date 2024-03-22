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

