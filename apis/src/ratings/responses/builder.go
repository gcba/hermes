package responses

import (
	"fmt"
	"net/http"

	"ratings/models"
	"ratings/parser"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type Meta struct {
	Code int `json:"code"`
	Message string `json:"message"`
}
