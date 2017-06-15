package models

import (
	"time"
)

type Message struct {
	ID        uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Message   string `sql:"type:text;not null"`
	Direction string `sql:"type:enum('in','out');not null"`
	RatingID  uint   `gorm:"not null"`

	CreatedAt time.Time `gorm:"not null"`
}
