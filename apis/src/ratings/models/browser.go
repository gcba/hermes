package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Browser struct {
	gorm.Model

	Name string `gorm:"size:15"`
}
