package models

import (
	"time"
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
