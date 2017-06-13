package models

import (
	"time"
)

type Device struct {
	ID           uint   `gorm:"primary_key"`
	Name         string `gorm:"size:30;not null"`
	ScreenWidth  int    `gorm:"not null"`
	ScreenHeight int    `gorm:"not null"`
	PPI          int
	Brand        Brand
	BrandID      int
	Platform     Platform
	PlatformID   int

	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time
}
