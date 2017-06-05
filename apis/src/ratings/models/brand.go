package models

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
    "ratings/models"
)

type Brand struct {
    gorm.Model

    Name string `gorm:"size:30"`
}