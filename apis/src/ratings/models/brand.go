package models

import (
	"time"
)

type Brand struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"size:30"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
