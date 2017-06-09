package models

import (
	"time"
)

type Platform struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"size:15"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
