package models

import (
	"github.com/jinzhu/gorm"
)

type Platform struct {
	gorm.Model

	Name string `gorm:"size:15"`
}
