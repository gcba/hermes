package models

import (
	"github.com/jinzhu/gorm"
)

type Brand struct {
	gorm.Model

	Name string `gorm:"size:30"`
}
