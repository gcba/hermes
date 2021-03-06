package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Device struct {
	ID           uint      `gorm:"primary_key;AUTO_INCREMENT"`
	Name         string    `gorm:"size:30;not null"`
	ScreenWidth  int       `gorm:"not null"`
	ScreenHeight int       `gorm:"not null"`
	PPI          *int      `gorm:"default:NULL"`
	BrandID      *uint     `gorm:"default:NULL"`
	PlatformID   uint      `gorm:"default:NULL"`
	AppUsers     []AppUser `gorm:"many2many:app_user_device;"`

	CreatedAt time.Time  `gorm:"not null;type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `gorm:"default:NULL"`
	DeletedAt *time.Time `gorm:"default:NULL"`
}

// GetDevice gets a device by name and brand id
func GetDevice(name string, db *gorm.DB) *gorm.DB {
	var result Device

	if isPostgres() {
		return db.Where("name ILIKE ?", name).First(&result)
	}

	return db.Where("name LIKE ?", name).First(&result)
}

// CreateDevice creates a new device
func CreateDevice(device *Device, db *gorm.DB) *gorm.DB {
	return db.Create(device)
}
