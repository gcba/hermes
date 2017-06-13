package models

import (
	"time"
)

type Platform struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"size:15;not null"`
	Key  string `gorm:"type:char(32);not null"`

	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time
}
