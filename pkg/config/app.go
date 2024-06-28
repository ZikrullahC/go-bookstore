package config

import (
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	// Gorm kullanarak MySQL bağlantısı açın
	d, err := gorm.Open("mysql", "root:ziko61ziko61@tcp(127.0.0.1:3306)/go-bookstore-database?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
