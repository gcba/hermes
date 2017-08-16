package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type App struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"size:30;not null"`
	Type string `gorm:"type:char;not null"`
	Key  string `gorm:"type:char(32);not null"`

	CreatedAt time.Time `gorm:"not null;type:datetime;default:CURRENT_TIMESTAMP"`
}

// GetApp gets an app by key
func GetApp(key string, db *gorm.DB) *gorm.DB {
	var result App

	return db.Where(&App{Key: key}).First(&result)
}
