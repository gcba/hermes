package models

import (
	"github.com/jinzhu/gorm"
)

type Message struct {
	gorm.Model

	Message   `sql:"type:text"`
	Direction `sql:"type:enum('in','out')`
	RatingID  int
}
