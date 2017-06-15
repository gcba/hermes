package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Device struct {
	ID           uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name         string `gorm:"size:30;not null"`
	ScreenWidth  int    `gorm:"not null"`
	ScreenHeight int    `gorm:"not null"`
	PPI          int
	Brand        Brand
	BrandID      uint
	Platform     Platform
	PlatformID   uint

	CreatedAt time.Time `gorm:"not null"`
}

// GetDevice gets a device by name and brand id
func GetDevice(name string, brand uint, db *gorm.DB) (Device, error) {
	var result Device

	query := "SELECT id, name FROM devices WHERE name LIKE ? AND brand_id = ?" // TODO: Set ILIKE

	if err := db.Raw(query, name, brand).Scan(&result).Error; err != nil {
		return Device{}, err
	}

	return result, nil
}
