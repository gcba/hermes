package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type Rating struct {
	ID              uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Rating          int8   `gorm:"not null"`
	Description     string `gorm:"size:30;DEFAULT:NULL"`
	AppVersion      string `gorm:"size:15;DEFAULT:NULL"`
	BrowserVersion  string `gorm:"size:15;DEFAULT:NULL"`
	PlatformVersion string `gorm:"size:15;DEFAULT:NULL"`
	HasMessage      bool   `gorm:"not null;DEFAULT:false"`
	AppID           uint   `gorm:"not null"`
	RangeID         uint   `gorm:"not null"`
	PlatformID      uint   `gorm:"DEFAULT:NULL"`
	DeviceID        uint   `gorm:"DEFAULT:NULL"`
	AppUserID       uint   `gorm:"column:appuser_id;DEFAULT:NULL"`
	BrowserID       uint   `gorm:"DEFAULT:NULL"`

	CreatedAt time.Time   `gorm:"not null;type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt pq.NullTime `gorm:"default:NULL"`
	DeletedAt pq.NullTime `gorm:"default:NULL"`
}

// CreateRating creates a new rating
func CreateRating(rating *Rating, db *gorm.DB) *gorm.DB {
	return db.Create(rating)
}
