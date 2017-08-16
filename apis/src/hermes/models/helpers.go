package models

import (
	"reflect"

	"github.com/jinzhu/gorm"
)

func isPostgres(db *gorm.DB) bool {
	driver := reflect.ValueOf(db.DB().Driver())

	if driver.Type().String() == "*pq.Driver" {
		return true
	}

	return false
}
