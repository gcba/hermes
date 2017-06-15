package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type AppUser struct {
	ID     uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name   string `gorm:"size:70;not null"`
	Email  string `gorm:"size:100;not null"`
	MiBAID uint   `gorm:"column:miba_id;not null"`

	CreatedAt time.Time `gorm:"not null"`
}

// TableName sets AppUser's table name to be `appuser`
func (AppUser) TableName() string {
	return "appusers"
}

// GetAppUser gets an app user by email addreess
func GetAppUser(mibaID uint, db *gorm.DB) (AppUser, error) {
	var result AppUser

	if err := db.Where(&AppUser{MiBAID: mibaID}).First(&result).Error; err != nil {
		return AppUser{}, err
	}

	return result, nil
}
