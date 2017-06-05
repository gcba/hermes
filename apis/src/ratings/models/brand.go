package models

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Brand struct {
    gorm.Model

    Name string `gorm:"size:30"`
}