package models

import (
	"time"
)

type Message struct {
	ID        uint   `gorm:"primary_key"`
	Message   string `sql:"type:text"`
	Direction string `sql:"type:enum('in','out')"`
	RatingID  int

	CreatedAt time.Time
	UpdatedAt time.Time
}
