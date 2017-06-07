package models

import (
	"github.com/jinzhu/gorm"
)

type Message struct {
	gorm.Model

	Message   string `sql:"type:text"`
	Direction string `sql:"type:enum('in','out')"`
	RatingID  int
}
