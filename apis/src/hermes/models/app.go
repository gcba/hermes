package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type App struct {
	ID       uint      `gorm:"primary_key"`
	Name     string    `gorm:"size:30;not null"`
	Type     string    `gorm:"type:char;not null"`
	Key      string    `gorm:"type:char(32);not null"`
	AppUsers []AppUser `gorm:"many2many:app_user_app;"`

	CreatedAt time.Time   `gorm:"not null;type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt pq.NullTime `gorm:"default:NULL"`
	UpdatedBy *uint       `gorm:"default:NULL"`
	DeletedAt pq.NullTime `gorm:"default:NULL"`
}

// GetApp gets an app by key
func GetApp(key string, db *gorm.DB) *gorm.DB {
	var result App

	return db.Where(&App{Key: key}).First(&result)
}
