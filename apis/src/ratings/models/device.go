package models

import (
	"github.com/jinzhu/gorm"
)

type Device struct {
	gorm.Model

	Name         string `gorm:"size:30"`
	ScreenWidth  int
	ScreenHeight int
	PPI          int
	Brand        Brand
	BrandID      int
}
