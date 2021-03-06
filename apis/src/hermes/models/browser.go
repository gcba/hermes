package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Browser struct {
	ID   uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name string `gorm:"size:15;not null"`

	CreatedAt time.Time  `gorm:"not null;type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `gorm:"default:NULL"`
	DeletedAt *time.Time `gorm:"default:NULL"`
}

// GetBrowser gets a browser by name
func GetBrowser(name string, db *gorm.DB) *gorm.DB {
	var result Browser

	if isPostgres() {
		return db.Where("name ILIKE ?", name).First(&result)
	}

	return db.Where("name LIKE ?", name).First(&result)
}

// CreateBrowser creates a new browser
func CreateBrowser(browser *Browser, db *gorm.DB) *gorm.DB {
	return db.Create(browser)
}
