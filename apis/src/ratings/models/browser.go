package models

import (
	"time"
)

type Browser struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"size:15;not null"`

	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time
}
