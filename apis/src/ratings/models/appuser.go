package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type AppUser struct {
	gorm.Model

	Name  string `gorm:"size:70"`
	Email string `gorm:"size:100"`
}

// set AppUser's table name to be `appuser`
func (AppUser) TableName() string {
	return "appuser"
}
