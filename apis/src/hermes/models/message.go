package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type Message struct {
	ID          uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Message     string `sql:"type:text;not null"`
	Direction   string `sql:"type:enum('in','out');not null"`
	Status      int    `gorm:"not null;DEFAULT:0"`
	TransportID *uint  `gorm:"DEFAULT:NULL"`
	RatingID    uint   `gorm:"not null"`

	CreatedAt time.Time   `gorm:"not null;type:datetime;default:CURRENT_TIMESTAMP"`
	CreatedBy *uint       `gorm:"default:NULL"`
	UpdatedAt pq.NullTime `gorm:"default:NULL"`
}

// CreateMessage creates a new message
func CreateMessage(message *Message, db *gorm.DB) *gorm.DB {
	return db.Create(message)
}
