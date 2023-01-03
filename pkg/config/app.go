package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	database *gorm.DB
)

func Connect() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	database = db
}

func GetDB() *gorm.DB {

	return database
}
