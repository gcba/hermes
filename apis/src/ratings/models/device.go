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
	PPI          *int   `gorm:"DEFAULT:NULL"`
	Brand        *Brand
	BrandID      *uint `gorm:"DEFAULT:NULL"`
	Platform     Platform
	PlatformID   uint `gorm:"DEFAULT:NULL"`

	CreatedAt time.Time `gorm:"not null;type:datetime;default:CURRENT_TIMESTAMP"`
}

// GetDevice gets a device by name and brand id
func GetDevice(name string, brand *Brand, db *gorm.DB) *gorm.DB {
	if brand != nil {
		return getDeviceByNameAndBrand(name, *brand, db)
	}

	return getDeviceByName(name, db)
}

// CreateDevice creates a new device
func CreateDevice(device *Device, db *gorm.DB) *gorm.DB {
	return db.Create(device)
}

func getDeviceByNameAndBrand(name string, brand Brand, db *gorm.DB) *gorm.DB {
	var result Device

	if isPostgres(db) {
		return db.Where("name ILIKE ? AND brand_id = ?", name, brand.ID).First(&result)
	}

	return db.Where("name LIKE ? AND brand_id = ?", name, brand).First(&result)
}

func getDeviceByName(name string, db *gorm.DB) *gorm.DB {
	var result Device

	if isPostgres(db) {
		return db.Where("name ILIKE ?", name).First(&result)
	}

	return db.Where("name LIKE ?", name).First(&result)
}
