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
	PPI          int    `gorm:"DEFAULT:NULL"`
	Brand        *Brand
	BrandID      uint `gorm:"DEFAULT:NULL"`
	Platform     *Platform
	PlatformID   uint `gorm:"DEFAULT:NULL"`

	CreatedAt time.Time `gorm:"not null"`
}

// GetDevice gets a device by name and brand id
func GetDevice(name string, brand uint, db *gorm.DB) *gorm.DB {
	var result Device
	var query string

	if isPostgres(db) {
		query = "SELECT id FROM devices WHERE name ILIKE ? AND brand_id = ?"
	} else {
		query = "SELECT id FROM devices WHERE name LIKE ? AND brand_id = ?"
	}

	return db.Raw(query, name, brand).Scan(&result)
}

// CreateDevice creates a new device
func CreateDevice(device *Device, db *gorm.DB) *gorm.DB {
	return db.Create(device)
}
