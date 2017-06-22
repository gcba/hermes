package models

import (
	"reflect"
	"time"

	"github.com/jinzhu/gorm"
)

type Brand struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"size:30;not null"`

	CreatedAt time.Time `gorm:"not null"`
}

// GetBrand gets a brand by name
func GetBrand(name string, db *gorm.DB) *gorm.DB {
	var result Brand
	var query string

	if isPostgres(db) {
		query = "SELECT id FROM brands WHERE name ILIKE ?"
	} else {
		query = "SELECT id FROM brands WHERE name LIKE ?"
	}

	return db.Raw(query, name).Scan(&result)
}

// CreateBrand creates a new brand
func CreateBrand(brand *Brand, db *gorm.DB) *gorm.DB {
	return db.Create(brand)
}

func isPostgres(db *gorm.DB) bool {
	driver := reflect.ValueOf(db.DB().Driver())

	if driver.Type().String() == "*pq.Driver" {
		return true
	}

	return false
}
