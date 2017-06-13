package models

import (
	"time"
)

type Brand struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"size:30;not null"`

	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time
}
