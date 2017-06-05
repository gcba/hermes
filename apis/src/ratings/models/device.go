package models

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Device struct {
    gorm.Model

    Name string `gorm:"size:30"`
    ScreenWidth int
    ScreenHeight int
    PPI int
    BrandID int
}