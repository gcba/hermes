package schema

import (
	"context"
	"fmt"
	"strings"

	"hermes/models"

	"github.com/fatih/structs"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/gorm"
	"github.com/neelance/graphql-go/errors"
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

	StatsError struct {
		*errors.QueryError
		Code int
	}

	Resolver struct{}
)

func (r *Resolver) Count(context context.Context, args arguments) (int32, error) {
	var total int32

	if db, castOk := context.Value(DB).(*gorm.DB); castOk {
		operator := args.Field.resolveOperator()
		count := fmt.Sprintf("COUNT(%s)", args.Field.Name)
		model := args.Field.getModel(db)
		query := db.Model(model).Select(count)
		entity := args.Field.getEntity()

		if model == nil {
			return total, invalidTableError(entity.Table)
		}

		if !fieldExists(entity.Field, structs.Names(model)) {
			return total, invalidFieldError(entity.Field)
		}

		if value := args.Field.getValue(); value != nil {
			query = query.Where(fmt.Sprintf("%s %s ?", args.Field.Name, operator), value)
		}

		query = args.attachAND(query)
		query = args.attachOR(query)
		query = query.Count(&total)

		errorList := query.GetErrors()

		if !(len(errorList) > 0 || query.Error != nil || query.Value == nil) {
			return total, nil
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
			return total, badRequestError("Average does not need a value on the main field")
		}

		average := fmt.Sprintf("AVG(%s)", args.Field.Name)
		model := args.Field.getModel(db)
		query := db.Model(model)
		entity := args.Field.getEntity()

		if model == nil {
			return total, invalidTableError(entity.Table)
		}

		if !fieldExists(entity.Field, structs.Names(model)) {
			return total, invalidFieldError(entity.Field)
		}

		query = query.Select(average)
		query = args.attachAND(query)
		query = args.attachOR(query)

		err := query.Row().Scan(&total)
		errorList := query.GetErrors()

		if !(len(errorList) > 0 || err != nil || query.Value == nil) {
			return total, nil
		} else if query.Error != nil {
			return total, queryError(err)
		}

		return total, databaseError()
	}

	return total, connectionError()
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

	return ""
}

func fieldExists(field string, fields []string) bool {
	for _, item := range fields {
		if item == strcase.ToCamel(field) {
			return true
		}
	}

	return false
}
