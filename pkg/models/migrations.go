package models

import (
	"github.com/matheusosorio/orm-server/pkg/config"
	"gorm.io/gorm"
)

func init() {
	config.Connect()
	db = config.GetDB()
	autoMigrations(db)
}

func autoMigrations(db *gorm.DB) {
	db.AutoMigrate(&Book{})
}
