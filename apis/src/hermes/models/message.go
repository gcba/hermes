package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Message struct {
	ID          uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Message     string `sql:"type:text;not null"`
	Direction   string `sql:"type:enum('in','out');not null"`
	Status      int    `gorm:"not null;DEFAULT:0"`
	RatingID    uint   `gorm:"not null"`
	TransportID *uint  `gorm:"DEFAULT:NULL"`

	CreatedAt time.Time `gorm:"not null;type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"-"`
}

// CreateMessage creates a new message
func CreateMessage(message *Message, db *gorm.DB) *gorm.DB {
	return db.Create(message)
}
