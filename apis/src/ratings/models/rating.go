package models

import (
	"time"
)

type Rating struct {
	ID              uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Rating          int8   `gorm:"not null"`
	Description     string `gorm:"size:30"`
	AppVersion      string `gorm:"size:15"`
	PlatformVersion string `gorm:"size:15"`
	BrowserVersion  string `gorm:"size:15"`
	HasMessage      bool   `gorm:"not null;DEFAULT:false"`
	AppID           int    `gorm:"not null"`
	RangeID         int    `gorm:"not null"`
	AppUserID       int    `gorm:"column:appuser_id"`
	PlatformID      int
	DeviceID        int
	BrowserID       int

	CreatedAt time.Time `gorm:"not null"`
}
