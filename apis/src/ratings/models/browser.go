package models

import (
	"github.com/jinzhu/gorm"
)

type Browser struct {
	gorm.Model

	Name string `gorm:"size:15"`
}
