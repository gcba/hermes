package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Browser struct {
	ID   uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name string `gorm:"size:15;not null"`

	CreatedAt time.Time `gorm:"not null"`
}

// GetBrowser gets a browser by name
func GetBrowser(name string, db *gorm.DB) *gorm.DB {
	var result Browser

	query := "SELECT id FROM browsers WHERE name LIKE ?" // TODO: Set ILIKE

	return db.Raw(query, name).Scan(&result)
}

// CreateBrowser creates a new browser
func CreateBrowser(browser *Browser, db *gorm.DB) *gorm.DB {
	return db.Create(browser)
}
