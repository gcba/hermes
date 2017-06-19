package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Range struct {
	ID    uint   `gorm:"primary_key;AUTO_INCREMENT"`
	From  uint8  `gorm:"not null"`
	To    uint8  `gorm:"not null"`
	Key   string `gorm:"type:char(32);not null"`
	AppID uint   `gorm:"not null"`

	CreatedAt time.Time `gorm:"not null"`
}

// GetRange gets a range by key
func GetRange(key string, db *gorm.DB) *gorm.DB {
	var result Range

	return db.Where(&Range{Key: key}).First(&result)
}
