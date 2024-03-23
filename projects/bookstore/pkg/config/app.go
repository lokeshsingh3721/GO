package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var(
	db * gorm.DB
)

func Connect (){
	d,err:= gorm.Open("mysql","avnadmin:AVNS_WqpMQ0ua4-Tg90ioU5K@mysql-156a842-study3721-9d82.a.aivencloud.com:15607/defaultdb?ssl-mode=REQUIRED")
	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}