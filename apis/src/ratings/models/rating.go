package models

import (
	"github.com/jinzhu/gorm"
)

type Rating struct {
	gorm.Model

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
}
