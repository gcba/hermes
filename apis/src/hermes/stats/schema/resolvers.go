package schema

import (
	"context"
	"errors"
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
			r.count(db, args, query, &total)

			return total, nil
		}

		return total, errors.New("Could not get value from database")
	}

	return total, errors.New("Could not connect to database")
}

func (r *Resolver) Average(context context.Context, args arguments) (float64, error) {
	var total float64

	if db, castOk := context.Value(DB).(*gorm.DB); castOk {
		entity := args.Field.getEntity()
		average := fmt.Sprintf("AVG(%s)", entity.Field)
		query := args.Field.query(db).Select(average)
		errorList := query.GetErrors()

		if !(len(errorList) > 0 || query.Error != nil || query.Value == nil) {
			r.average(db, args, query, &total)

			return total, nil
		}

		return total, errors.New("Could not get value from database")
	}

	return total, errors.New("Could not connect to database")
}

func (r *Resolver) count(db *gorm.DB, args arguments, query *gorm.DB, total *int32) {
	if args.And == nil && args.Or == nil {
		query.Count(total)
	} else if args.And != nil {
		*total = 9 // Placeholder value
	} else if args.Or != nil {
		for _, item := range *args.Or {
			var subtotal int32

			subquery := item.query(db)
			errorList := subquery.GetErrors()

			if !(len(errorList) > 0 || subquery.Error != nil || subquery.Value == nil) {
				subquery.Count(&subtotal)

				*total += subtotal
			}
		}
	}
}

func (r *Resolver) average(db *gorm.DB, args arguments, query *gorm.DB, total *float64) {
	if args.And == nil && args.Or == nil {
		query.Row().Scan(total)
	} else if args.And != nil {
		*total = 8 // Placeholder value
	} else if args.Or != nil {
		*total = 9 // Placeholder value
	}
}

func (f *field) query(db *gorm.DB) *gorm.DB {
	operator := f.resolveOperator()
	entity := f.getEntity()
	value := f.getValue()
	where := fmt.Sprintf("%s %s ?", entity.Field, operator)

	switch entity.Table {
	case "apps":
		return db.Model(&models.Rating{}).Where(where, value)
	case "appusers":
		return db.Model(&models.AppUser{}).Where(where, value)
	case "brands":
		return db.Model(&models.Brand{}).Where(where, value)
	case "browsers":
		return db.Model(&models.Browser{}).Where(where, value)
	case "devices":
		return db.Model(&models.Device{}).Where(where, value)
	case "messages":
		return db.Model(&models.Message{}).Where(where, value)
	case "platforms":
		return db.Model(&models.Platform{}).Where(where, value)
	case "ranges":
		return db.Model(&models.Range{}).Where(where, value)
	case "ratings":
		fallthrough
	default:
		return db.Model(&models.Rating{}).Where(where, value)
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
