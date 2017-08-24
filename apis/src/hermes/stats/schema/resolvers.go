package schema

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"hermes/models"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type (
	entity struct {
		Table string
		Field string
	}

	field struct {
		Name string
		Eq   *Value
	}

	arguments struct {
		Field field
		Or    *[]field
		And   *[]field
	}

	Resolver struct{}
)

func errorResponse() error {
	return echo.NewHTTPError(http.StatusInternalServerError)
}

func (r *Resolver) Count(context context.Context, args arguments) (int32, error) {
	var total int32

	if db, castOk := context.Value(DB).(*gorm.DB); castOk {
		query := args.Field.query(db)
		errorList := query.GetErrors()

		if !(len(errorList) > 0 || query.Error != nil || query.Value == nil) {
			var count int32

			query.Count(&count)

			total += count
		}
	}

	// TODO: Handle non existent db

	return total, nil
}

func (r *Resolver) Average(context context.Context, args arguments) (float64, error) {
	// TODO: Implement

	return 0, nil
}

func (f *field) query(db *gorm.DB) *gorm.DB {
	operator := f.resolveOperator()
	entity := f.getEntity()
	value := f.getValue()
	where := fmt.Sprintf("%s %s ?", entity.Field, operator)

	switch entity.Table {
	case "apps":
		return db.Where(where, value).Find(&[]models.App{})
	case "appusers":
		return db.Where(where, value).Find(&[]models.AppUser{})
	case "brands":
		return db.Where(where, value).Find(&[]models.Brand{})
	case "browsers":
		return db.Where(where, value).Find(&[]models.Browser{})
	case "devices":
		return db.Where(where, value).Find(&[]models.Device{})
	case "messages":
		return db.Where(where, value).Find(&[]models.Message{})
	case "platforms":
		return db.Where(where, value).Find(&[]models.Platform{})
	case "ranges":
		return db.Where(where, value).Find(&[]models.Range{})
	case "ratings":
		fallthrough
	default:
		return db.Where(where, value).Find(&[]models.Rating{})
	}
}

func (f *field) getEntity() entity {
	splitField := strings.Split(f.Name, ".")

	return entity{Table: splitField[0], Field: splitField[1]}
}

func (f *field) getValue() interface{} {
	if f.Eq != nil {
		if f.Eq.String != nil {
			return f.Eq.String
		} else if f.Eq.Int != nil {
			return f.Eq.Int
		} else if f.Eq.Float != nil {
			return f.Eq.Float
		} else if f.Eq.Bool != nil {
			return f.Eq.Bool
		}
	}

	return nil
}

func (f *field) resolveOperator() string {
	value := f.getValue()

	if f.Eq != nil {
		switch value.(type) {
		case string:
			return "LIKE"
		default:
			return "="
		}
	}

	return "="
}
