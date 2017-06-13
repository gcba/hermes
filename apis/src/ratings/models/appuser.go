package models

import (
	"time"
)

type AppUser struct {
	ID     uint   `gorm:"primary_key"`
	Name   string `gorm:"size:70"`
	Email  string `gorm:"size:100"`
	MiBAID uint   `gorm:"column:miba_id"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName sets AppUser's table name to be `appuser`
func (AppUser) TableName() string {
	return "appuser"
}
