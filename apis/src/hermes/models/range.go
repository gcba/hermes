package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Range struct {
	ID    uint   `gorm:"primary_key;AUTO_INCREMENT"`
	From  int8   `gorm:"not null"`
	To    int8   `gorm:"not null"`
	Key   string `gorm:"type:char(32);not null"`
	AppID uint   `gorm:"not null"`

	CreatedAt time.Time `gorm:"not null;type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"-"`
	DeletedAt time.Time `gorm:"-"`
}

// GetRange gets a range by key
func GetRange(key string, db *gorm.DB) *gorm.DB {
	var result Range

	return db.Where(&Range{Key: key}).First(&result)
}
