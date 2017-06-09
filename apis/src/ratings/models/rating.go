package models

import (
	"time"
)

type Rating struct {
	ID              uint `gorm:"primary_key"`
	Rating          int8
	Description     string `gorm:"size:30"`
	AppVersion      string `gorm:"size:15"`
	PlatformVersion string `gorm:"size:15"`
	BrowserVersion  string `gorm:"size:15"`
	HasMessage      bool
	AppID           int
	AppUserID       int `gorm:"column:appuser_id"`
	PlatformID      int
	DeviceID        int
	BrowserID       int

	CreatedAt time.Time
	UpdatedAt time.Time
}
