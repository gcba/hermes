package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Platform struct {
	ID   uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name string `gorm:"size:15;not null"`
	Key  string `gorm:"type:char(32);not null"`

	CreatedAt time.Time `gorm:"not null;type:datetime;default:CURRENT_TIMESTAMP"`
}

// GetPlatform gets a platform by key
func GetPlatform(key string, db *gorm.DB) *gorm.DB {
	var result Platform

	return db.Where(&Platform{Key: key}).First(&result)
}
