package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type AppUser struct {
	ID        uint       `gorm:"primary_key;AUTO_INCREMENT"`
	Name      string     `gorm:"size:70;not null"`
	Email     *string    `gorm:"size:100;default:NULL"`
	MiBAID    *string    `gorm:"column:miba_id;type:char(36);default:NULL"`
	Apps      []App      `gorm:"many2many:app_user_app;"`
	Platforms []Platform `gorm:"many2many:app_user_platform;"`
	Devices   []Device   `gorm:"many2many:app_user_device;"`

	CreatedAt time.Time  `gorm:"not null;type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `gorm:"default:NULL"`
	DeletedAt *time.Time `gorm:"default:NULL"`
}

// TableName sets AppUser's table name to be `appuser`
func (AppUser) TableName() string {
	return "appusers"
}

// GetAppUser gets an app user by MiBA id
func GetAppUser(mibaID string, db *gorm.DB) *gorm.DB { // TODO: Check actual mibaID type
	var result AppUser

	return db.Where(&AppUser{MiBAID: &mibaID}).First(&result)
}

// GetAppUserByEmail gets an app user by email
func GetAppUserByEmail(email string, db *gorm.DB) *gorm.DB {
	var result AppUser

	return db.Where(&AppUser{Email: &email}).First(&result)
}

// CreateAppUser creates a new app user
func CreateAppUser(appuser *AppUser, db *gorm.DB) *gorm.DB {
	return db.Create(appuser)
}
