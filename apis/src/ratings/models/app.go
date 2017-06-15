package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type App struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"size:30;not null"`
	Type rune   `gorm:"type:char;not null"`
	Key  string `gorm:"type:char(32);not null"`

	CreatedAt time.Time `gorm:"not null"`
}

// GetApp gets an app by key
func GetApp(key string, db *gorm.DB) (App, error) {
	var result App

	query := "SELECT id FROM apps WHERE key = ?"

	if err := db.Raw(query, key).Scan(&result).Error; err != nil {
		return App{}, err
	}

	return result, nil
}
