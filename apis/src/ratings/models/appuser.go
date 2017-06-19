package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type AppUser struct {
	ID     uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name   string `gorm:"size:70;not null"`
	Email  string `gorm:"size:100;not null"`
	MiBAID string `gorm:"column:miba_id;not null"`

	CreatedAt time.Time `gorm:"not null"`
}

// TableName sets AppUser's table name to be `appuser`
func (AppUser) TableName() string {
	return "appusers"
}

// GetAppUser gets an app user by MiBA id
func GetAppUser(mibaID string, db *gorm.DB) *gorm.DB { // TODO: Check actual mibaID type
	var result AppUser

	return db.Where(&AppUser{MiBAID: mibaID}).First(&result)
}

// CreateAppUser creates a new app user
func CreateAppUser(appuser *AppUser, db *gorm.DB) *gorm.DB {
	return db.Create(appuser)
}
