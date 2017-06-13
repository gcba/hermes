package models

import (
	"time"
)

type AppUser struct {
	ID     uint   `gorm:"primary_key"`
	Name   string `gorm:"size:70;not null"`
	Email  string `gorm:"size:100;not null"`
	MiBAID uint   `gorm:"column:miba_id;not null"`

	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time
}

// TableName sets AppUser's table name to be `appuser`
func (AppUser) TableName() string {
	return "appuser"
}
