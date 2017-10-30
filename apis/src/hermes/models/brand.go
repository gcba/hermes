package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Brand struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"size:30;not null"`

	CreatedAt time.Time `gorm:"not null;type:datetime;default:CURRENT_TIMESTAMP"`
}

// GetBrand gets a brand by name
func GetBrand(name string, db *gorm.DB) *gorm.DB {
	var result Brand

	if isPostgres() {
		return db.Where("name ILIKE ?", name).First(&result)
	}

	return db.Where("name LIKE ?", name).First(&result)
}

// CreateBrand creates a new brand
func CreateBrand(brand *Brand, db *gorm.DB) *gorm.DB {
	return db.Create(brand)
}
