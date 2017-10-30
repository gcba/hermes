package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type Platform struct {
	ID       uint      `gorm:"primary_key;AUTO_INCREMENT"`
	Name     string    `gorm:"size:15;not null"`
	Key      string    `gorm:"type:char(32);not null"`
	AppUsers []AppUser `gorm:"many2many:app_user_platform;"`

	CreatedAt time.Time   `gorm:"not null;type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt pq.NullTime `gorm:"default:NULL"`
	DeletedAt pq.NullTime `gorm:"default:NULL"`
}

// GetPlatform gets a platform by key
func GetPlatform(key string, db *gorm.DB) *gorm.DB {
	var result Platform

	return db.Where(&Platform{Key: key}).First(&result)
}
