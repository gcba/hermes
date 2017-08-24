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
	operator := args.Field.resolveOperator()
	where := fmt.Sprintf("%s %s ?", args.Field.Name, operator)

	query = query.Where(where, args.Field.getValue())

	if args.And != nil {
		for _, item := range *args.And {
			suboperator := item.resolveOperator()
			where := fmt.Sprintf("%s %s ?", item.Name, suboperator)

			query = query.Where(where, item.getValue())
		}
	}

	if args.Or != nil {
		for _, item := range *args.Or {
			suboperator := item.resolveOperator()
			where := fmt.Sprintf("%s %s ?", item.Name, suboperator)

			query = query.Or(where, item.getValue())
		}
	}

	query.Count(total)
}

func (r *Resolver) average(db *gorm.DB, args arguments, query *gorm.DB, total *float64) {
	operator := args.Field.resolveOperator()
	entity := args.Field.getEntity()
	value := args.Field.getValue()
	where := fmt.Sprintf("%s %s ?", entity.Field, operator)

	if args.And == nil && args.Or == nil {
		query.Where(where, value).Row().Scan(total)
	} else if args.And != nil {
		*total = 8 // Placeholder value
	} else if args.Or != nil {
		/*
			for _, item := range *args.Or {

					var subtotal int32

					entity := args.Field.getEntity()
					average := fmt.Sprintf("AVG(%s)", entity.Field)
					subquery := item.query(db).Select(average)
					errorList := subquery.GetErrors()

					if !(len(errorList) > 0 || subquery.Error != nil || subquery.Value == nil) {
						subquery.Row().Scan(&subtotal)

						*total += subtotal
					}

			}
		*/
	}
}

func (f *field) query(db *gorm.DB) *gorm.DB {
	entity := f.getEntity()

	switch entity.Table {
	case "apps":
		return db.Model(&models.Rating{})
	case "appusers":
		return db.Model(&models.AppUser{})
	case "brands":
		return db.Model(&models.Brand{})
	case "browsers":
		return db.Model(&models.Browser{})
	case "devices":
		return db.Model(&models.Device{})
	case "messages":
		return db.Model(&models.Message{})
	case "platforms":
		return db.Model(&models.Platform{})
	case "ranges":
		return db.Model(&models.Range{})
	case "ratings":
		fallthrough
	default:
		return db.Model(&models.Rating{})
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
