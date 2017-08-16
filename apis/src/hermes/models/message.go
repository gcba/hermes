package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Message struct {
	ID          uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Message     string `sql:"type:text;not null"`
	Direction   string `sql:"type:enum('in','out');not null"`
	Notified    bool   `gorm:"not null;DEFAULT:false"`
	RatingID    uint   `gorm:"not null"`
	TransportID *uint  `gorm:"DEFAULT:NULL"`

	CreatedAt time.Time `gorm:"not null;type:datetime;default:CURRENT_TIMESTAMP"`
}

// CreateMessage creates a new message
func CreateMessage(message *Message, db *gorm.DB) *gorm.DB {
	return db.Create(message)
}
