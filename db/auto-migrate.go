package db

import (
	"github.com/channyeintun/teacher-booking-system/models"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{})
}
