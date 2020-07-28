package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go-img/model"
)

var db *gorm.DB

func InitDb() {
	db, _ := gorm.Open("mysql", "root:@/go_img?charset=utf8&parseTime=True&loc=Local")

	db.LogMode(true)

	db.AutoMigrate(&model.Image{})

}

func Get() *gorm.DB {
	return db
}

func Close() {
	db.Close()
}
