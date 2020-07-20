package config


import (
	"awesomeProject/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open("mysql", "root:baiwa1234@(192.168.1.142:3306)/testk?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&entity.User{})
	DB = db
}