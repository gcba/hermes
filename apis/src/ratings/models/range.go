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
func GetRange(key string, db *gorm.DB) (Range, error) {
	var result Range

	query := "SELECT id FROM ranges WHERE key = ?"

	if err := db.Raw(query, key).Scan(&result).Error; err != nil {
		return Range{}, err
	}

	return result, nil
}
