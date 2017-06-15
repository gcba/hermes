package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Brand struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"size:30;not null"`

	CreatedAt time.Time `gorm:"not null"`
}

// GetBrand gets a brand by name
func GetBrand(name string, db *gorm.DB) *DB {
	var result Brand

	query := "SELECT id FROM brands WHERE name LIKE ?" // TODO: Set ILIKE

	return db.Raw(query, name).Scan(&result)∂
}

// CreateBrand creates a new brand
func CreateBrand(brand *Brand, db *gorm.DB) *DB {
	return db.Create(brand)
}
