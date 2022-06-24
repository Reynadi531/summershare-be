package database

import (
	"gorm.io/gorm"
	"summershare/internal/entities"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Post{})
}
