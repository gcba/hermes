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
		Name      string
		Operator  string
		Value     Value
		Next      *operation
		Condition *string
	}

	operation struct {
		Condition string
		Field     field
	}

	Resolver struct{}
)

func errorResponse() error {
	return echo.NewHTTPError(http.StatusInternalServerError)
}

func (r *Resolver) Count(context context.Context, args struct{ Field field }) (int32, error) {
	var total int32

	if db, castOk := context.Value(DB).(*gorm.DB); castOk {
		fields := args.Field.flatten([]*field{})
		queries := []*gorm.DB{}

		for index, field := range fields {
			query := field.query(db)
			errorList := query.GetErrors()

			queries = append(queries, query)

			if !(len(errorList) > 0 || query.Error != nil || query.Value == nil) {
				if field.Condition != nil {
					switch *field.Condition {
					case "OR":
						var count int32

						query.Count(&count)

						total += count
					case "AND":
						// TODO: Implement reducer for AND queries
					}
				} else if index == 0 {
					var count int32

					query.Count(&count)

					total += count
				}
			}
			// TODO: Handle db error
		}
	}

	// TODO: Handle non existent db

	return total, nil
}

func (r *Resolver) Average(context context.Context, args struct{ Field field }) (float64, error) {
	// TODO: Implement

	return 0, nil
}

func (f *field) flatten(buffer []*field) []*field {
	if f.Next != nil {
		f.Next.Field.Condition = &f.Next.Condition
		f.Next.Field.flatten(buffer)
	} else {
		buffer = append(buffer, f)
	}

	return buffer
}

func (f *field) query(db *gorm.DB) *gorm.DB {
	operator := f.resolveOperator(f.Value)
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
	if f.Value.String != nil {
		return f.Value.String
	} else if f.Value.Int != nil {
		return f.Value.Int
	} else if f.Value.Float != nil {
		return f.Value.Float
	} else if f.Value.Bool != nil {
		return f.Value.Bool
	}

	return nil
}

func (f *field) resolveOperator(value interface{}) string {
	switch f.Operator {
	case "EQ":
		return f.resolveEQOperator(value)
	}

	return f.Operator
}

func (f *field) resolveEQOperator(value interface{}) string {
	switch value.(type) {
	case string:
		return "LIKE"
	default:
		return "="
	}
}
