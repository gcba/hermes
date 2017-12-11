package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Rating struct {
	ID              uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Rating          int8   `gorm:"not null"`
	Description     string `gorm:"size:30;default:NULL"`
	AppVersion      string `gorm:"size:15;default:NULL"`
	BrowserVersion  string `gorm:"size:15;default:NULL"`
	PlatformVersion string `gorm:"size:15;default:NULL"`
	HasMessage      bool   `gorm:"not null;default:false"`
	AppID           uint   `gorm:"not null"`
	RangeID         uint   `gorm:"not null"`
	PlatformID      uint   `gorm:"default:NULL"`
	DeviceID        uint   `gorm:"default:NULL"`
	AppUserID       uint   `gorm:"column:appuser_id;default:NULL"`
	BrowserID       uint   `gorm:"default:NULL"`

	CreatedAt time.Time  `gorm:"not null;type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `gorm:"default:NULL"`
	DeletedAt *time.Time `gorm:"default:NULL"`
}

// CreateRating creates a new rating
func CreateRating(rating *Rating, db *gorm.DB) *gorm.DB {
	return db.Create(rating)
}
