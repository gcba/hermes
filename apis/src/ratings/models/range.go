package models

import (
	"time"
)

type Range struct {
	ID    uint   `gorm:"primary_key;AUTO_INCREMENT"`
	From  uint8  `gorm:"not null"`
	To    uint8  `gorm:"not null"`
	Key   string `gorm:"type:char(32);not null"`
	App   App
	AppID uint `gorm:"not null"`

	CreatedAt time.Time `gorm:"not null"`
}
