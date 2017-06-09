package models

import (
	"time"
)

type Device struct {
	ID           uint   `gorm:"primary_key"`
	Name         string `gorm:"size:30"`
	ScreenWidth  int
	ScreenHeight int
	PPI          int
	Brand        Brand
	BrandID      int
	Platform     Platform
	PlatformID   int

	CreatedAt time.Time
	UpdatedAt time.Time
}
