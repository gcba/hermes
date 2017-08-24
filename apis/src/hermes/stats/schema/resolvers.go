package schema

import (
	"context"
	"fmt"
	"strings"

	"hermes/models"

	"github.com/jinzhu/gorm"
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

func (r *Resolver) Count(context context.Context, args arguments) (int32, error) {
	var total int32

	if db, castOk := context.Value(DB).(*gorm.DB); castOk {
		operator := args.Field.resolveOperator()
		where := fmt.Sprintf("%s %s ?", args.Field.Name, operator)
		query := args.Field.getQuery(db).Where(where, args.Field.getValue())

		query = args.attachAND(query)
		query = args.attachOR(query)

		query.Count(&total)

		errorList := query.GetErrors()

		if !(len(errorList) > 0 || query.Error != nil || query.Value == nil) {
			return total, nil
		} else if query.Error != nil {
			return total, fmt.Errorf("Error getting value from database: %v", query.Error)
		}

		return total, fmt.Errorf("Could not get value from database")
	}

	return total, fmt.Errorf("Could not connect to database")
}

func (r *Resolver) Average(context context.Context, args arguments) (float64, error) {
	var total float64

	if db, castOk := context.Value(DB).(*gorm.DB); castOk {
		average := fmt.Sprintf("AVG(%s)", args.Field.Name)
		query := args.Field.getQuery(db).Select(average)

		query = args.attachAND(query)
		query = args.attachOR(query)

		query.Row().Scan(&total)

		errorList := query.GetErrors()

		if !(len(errorList) > 0 || query.Error != nil || query.Value == nil) {
			return total, nil
		} else if query.Error != nil {
			return total, fmt.Errorf("Error getting value from database: %v", query.Error)
		}

		return total, fmt.Errorf("Could not get value from database")
	}

	return total, fmt.Errorf("Could not connect to database")
}

func (a arguments) attachAND(query *gorm.DB) *gorm.DB {
	if a.And != nil {
		for _, item := range *a.And {
			operator := item.resolveOperator()
			where := fmt.Sprintf("%s %s ?", item.Name, operator)

			query = query.Where(where, item.getValue())
		}
	}

	return query
}

func (a arguments) attachOR(query *gorm.DB) *gorm.DB {
	if a.Or != nil {
		for _, item := range *a.Or {
			operator := item.resolveOperator()
			where := fmt.Sprintf("%s %s ?", item.Name, operator)

			query = query.Or(where, item.getValue())
		}
	}

	return query
}

func (f *field) getQuery(db *gorm.DB) *gorm.DB {
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
		return f.resolveValue(f.Eq)
	}

	return nil
}

func (f *field) resolveValue(value *Value) interface{} {
	if value.String != nil {
		return value.String
	} else if value.Int != nil {
		return value.Int
	} else if value.Float != nil {
		return value.Float
	} else if value.Bool != nil {
		return value.Bool
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
