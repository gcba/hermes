package models

import (
	"time"
)

type Browser struct {
	ID   uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name string `gorm:"size:15;not null"`

	CreatedAt time.Time `gorm:"not null"`
}
