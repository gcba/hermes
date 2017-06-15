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
func GetBrowser(name string, db *gorm.DB) (Browser, error) {
	var result Browser

	query := "SELECT id FROM browsers WHERE name LIKE ?" // TODO: Set ILIKE

	if err := db.Raw(query, name).Scan(&result).Error; err != nil {
		return Browser{}, err
	}

	return result, nil
}

// CreateBrowser creates a new browser
func CreateBrowser(browser *Browser, db *gorm.DB) error {
	if err := db.Create(browser).Error; err != nil {
		return err
	}

	return nil
}
