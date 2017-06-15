package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Platform struct {
	ID   uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name string `gorm:"size:15;not null"`
	Key  string `gorm:"type:char(32);not null"`

	CreatedAt time.Time `gorm:"not null"`
}

// GetPlatform gets a platform by key
func GetPlatform(name string, db *gorm.DB) (Platform, error) {
	var result Platform

	query := "SELECT id FROM platforms WHERE key = ?" // TODO: Set ILIKE

	if err := db.Raw(query, name).Scan(&result).Error; err != nil {
		return Platform{}, err
	}

	return result, nil
}
