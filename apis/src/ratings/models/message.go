package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Message struct {
	gorm.Model

	Message   `sql:"type:text"`
	Direction `sql:"type:enum('in','out')`
	RatingID  int
}
