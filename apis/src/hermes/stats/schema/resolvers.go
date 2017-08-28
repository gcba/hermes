package schema

import (
	"context"
	"fmt"
	"strings"

	"hermes/models"

	"github.com/fatih/structs"
	"github.com/jinzhu/gorm"
	"github.com/neelance/graphql-go/errors"
)

type (
	entity struct {
		Table string
		Field *string
	}

	field struct {
		Name  string
		Eq    *Value
		Ne    *Value
		Gt    *Value
		Lt    *Value
		Gte   *Value
		Lte   *Value
		Count *bool
	}

	arguments struct {
		Field field
		Or    *[]field
		And   *[]field
	}

	StatsError struct {
		*errors.QueryError
		Code int
	}

	Resolver struct{}
)

func (r *Resolver) Count(context context.Context, args arguments) (int32, error) {
	var total int32

	if db, castOk := context.Value(DB).(*gorm.DB); castOk {
		count := fmt.Sprintf("COUNT(%s)", args.Field.Name)
		model := args.Field.getModel(db)
		query := db.Model(model).Select(count)
		entity := args.Field.getEntity()

		if model == nil {
			return total, invalidTableError(entity.Table)
		}

		if entity.Field != nil && !fieldExists(*entity.Field, structs.Names(model)) {
			return total, invalidFieldError(*entity.Field)
		}

		if value := args.Field.getValue(); value != nil {
			operator := args.Field.resolveOperator()
			where := fmt.Sprintf("%s %s ?", args.Field.Name, operator)

			query = query.Where(where, value)
		}

		query = args.attachAND(query)
		query = args.attachOR(query)

		rows, err := query.Rows()
		errorList := query.GetErrors()
		accumulator := 0

		for rows.Next() {
			accumulator++
		}

		if accumulator > 1 {
			total = int32(accumulator)
		} else {
			query.Count(&total)
		}

		if !(len(errorList) > 0 || err != nil || query.Error != nil || query.Value == nil) {
			return total, nil
		} else if err != nil {
			return total, queryError(err)
		} else if query.Error != nil {
			return total, queryError(query.Error)
		}

		return total, databaseError()
	}

	return total, connectionError()
}

func (r *Resolver) Average(context context.Context, args arguments) (float64, error) {
	var total float64

	if db, castOk := context.Value(DB).(*gorm.DB); castOk {
		if value := args.Field.getValue(); value != nil {
			return total, badRequestError("Average does not accept a value on the main field")
		}

		average := fmt.Sprintf("AVG(%s)", args.Field.Name)
		model := args.Field.getModel(db)
		query := db.Model(model)
		entity := args.Field.getEntity()

		if model == nil {
			return total, invalidTableError(entity.Table)
		}

		if entity.Field == nil {
			return total, badRequestError("Average requires a field name")
		}

		if !fieldExists(*entity.Field, structs.Names(model)) {
			return total, invalidFieldError(*entity.Field)
		}

		query = query.Select(average)
		query = args.attachAND(query)
		query = args.attachOR(query)

		rows, err := query.Rows()
		errorList := query.GetErrors()
		numberOfRows := 0

		for rows.Next() {
			var avg float64

			rows.Scan(&avg)

			total += avg
			numberOfRows++
		}

		total = total / float64(numberOfRows)

		if !(len(errorList) > 0 || err != nil || query.Error != nil || query.Value == nil) {
			return total, nil
		} else if err != nil {
			return total, queryError(err)
		} else if query.Error != nil {
			return total, queryError(query.Error)
		}

		return total, databaseError()
	}

	return total, connectionError()
}

func (a arguments) attachAND(query *gorm.DB) *gorm.DB {
	if a.And != nil {
		for _, item := range *a.And {
			operator := item.resolveOperator()

			if item.Count != nil && *item.Count {
				having := fmt.Sprintf("COUNT(%s) %s ?", item.Name, operator)

				query = query.Group(item.Name).Having(having, item.getValue())
			} else {
				where := fmt.Sprintf("%s %s ?", item.Name, operator)

				query = query.Where(where, item.getValue())
			}
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

func (f *field) getModel(db *gorm.DB) interface{} {
	entity := f.getEntity()

	switch entity.Table {
	case "apps":
		return &models.App{}
	case "appusers":
		return &models.AppUser{}
	case "brands":
		return &models.Brand{}
	case "browsers":
		return &models.Browser{}
	case "devices":
		return &models.Device{}
	case "messages":
		return &models.Message{}
	case "platforms":
		return &models.Platform{}
	case "ranges":
		return &models.Range{}
	case "ratings":
		return &models.Rating{}
	default:
		return nil
	}
}

func (f *field) getEntity() entity {
	splitField := strings.Split(f.Name, ".")

	if len(splitField) < 2 {
		return entity{Table: splitField[0], Field: nil}
	}

	return entity{Table: splitField[0], Field: &splitField[1]}
}

func (f *field) getValue() interface{} {
	if f.Eq != nil {
		return f.resolveValue(f.Eq)
	} else if f.Ne != nil {
		return f.resolveValue(f.Ne)
	} else if f.Gt != nil {
		return f.resolveValue(f.Gt)
	} else if f.Lt != nil {
		return f.resolveValue(f.Lt)
	} else if f.Gte != nil {
		return f.resolveValue(f.Gte)
	} else if f.Lte != nil {
		return f.resolveValue(f.Lte)
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
		case string, *string:
			if isPostgres() {
				return "ILIKE"
			}

			return "LIKE"
		default:
			return "="
		}
	} else if f.Ne != nil {
		switch value.(type) {
		case string, *string:
			if isPostgres() {
				return "NOT ILIKE"
			}

			return "NOT LIKE"
		default:
			return "<>"
		}
	} else if f.Gt != nil {
		return ">"
	} else if f.Lt != nil {
		return "<"
	} else if f.Gte != nil {
		return ">="
	} else if f.Lte != nil {
		return "<="
	}

	return "="
}
