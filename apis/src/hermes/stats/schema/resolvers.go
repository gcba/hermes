package schema

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"hermes/models"

	"github.com/fatih/structs"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/gorm"
	graphqlErrors "github.com/neelance/graphql-go/errors"
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

	countResult struct {
		Count int32
	}

	averageResult struct {
		Average float64
	}

	StatsError struct {
		*graphqlErrors.QueryError
		Code int
	}

	Resolver struct{}
)

func (r *Resolver) Count(context context.Context, args arguments) (int32, error) {
	var result countResult

	if db, castOk := context.Value(DB).(*gorm.DB); castOk {
		model := args.Field.getModel(db)
		entity := args.Field.getEntity()

		if model == nil {
			return result.Count, invalidTableError(entity.Table)
		}

		if entity.Field != nil && !fieldExists(*entity.Field, structs.Names(model)) {
			return result.Count, invalidFieldError(*entity.Field)
		}

		count := fmt.Sprintf("COUNT(%s) AS Count", args.Field.Name)
		query := db.Select(count).Table(entity.Table)
		modelStruct := structs.New(model)

		if entity.Field != nil {
			fieldName := toCamelCase(*entity.Field)
			field := modelStruct.Field(fieldName)
			fieldKind := field.Kind()

			if operator := args.Field.resolveOperator(); operator != nil {
				if value := args.Field.getValue(); value != nil {
					valueType := reflect.TypeOf(value)

					if isValidType(fieldKind, valueType) {
						where := fmt.Sprintf("%s %s ?", args.Field.Name, *operator)

						query = query.Where(where, value)
					} else {
						return result.Count, invalidValueError()
					}
				} else {
					return result.Count, noValueError()
				}
			}
		}

		query = args.attachAND(query)
		query = args.attachOR(query)

		query.Scan(&result)

		errorList := query.GetErrors()

		if !(len(errorList) > 0 || query.Error != nil) {
			return result.Count, nil
		} else if query.Error != nil {
			return result.Count, queryError(query.Error)
		}

		return result.Count, databaseError()
	}

	return result.Count, connectionError()
}

func (r *Resolver) Average(context context.Context, args arguments) (float64, error) {
	var result averageResult

	if db, castOk := context.Value(DB).(*gorm.DB); castOk {
		model := args.Field.getModel(db)
		entity := args.Field.getEntity()

		if model == nil {
			return result.Average, invalidTableError(entity.Table)
		}

		average := fmt.Sprintf("AVG(%s) AS Average", args.Field.Name)
		query := db.Select(average).Table(entity.Table)

		if entity.Field == nil {
			return result.Average, badRequestError("Average requires a field name")
		}

		if !fieldExists(*entity.Field, structs.Names(model)) {
			return result.Average, invalidFieldError(*entity.Field)
		}

		modelStruct := structs.New(model)
		fieldName := toCamelCase(*entity.Field)
		field := modelStruct.Field(fieldName)
		fieldKind := field.Kind()

		if fieldKind == reflect.Ptr {
			fieldValueKind := reflect.TypeOf(field.Value()).Kind()

			if !isNumericKind(fieldValueKind) {
				return result.Average, invalidFieldError(*entity.Field)
			}
		} else if !isNumericKind(fieldKind) {
			return result.Average, invalidFieldError(*entity.Field)
		}

		if operator := args.Field.resolveOperator(); operator != nil {
			if value := args.Field.getValue(); value != nil {
				valueType := reflect.TypeOf(value)

				if isValidNumericType(valueType) {
					where := fmt.Sprintf("%s %s ?", args.Field.Name, *operator)

					query = query.Where(where, value)
				} else {
					return result.Average, invalidValueError()
				}
			} else {
				return result.Average, noValueError()
			}
		}

		query = args.attachAND(query)
		query = args.attachOR(query)

		query.Scan(&result)

		errorList := query.GetErrors()

		if !(len(errorList) > 0 || query.Error != nil) {
			return result.Average, nil
		} else if query.Error != nil {
			return result.Average, queryError(query.Error)
		}

		return result.Average, databaseError()
	}

	return result.Average, connectionError()
}

func (a arguments) attachAND(query *gorm.DB) *gorm.DB {
	if a.And != nil {
		for _, item := range *a.And {
			if operator := item.resolveOperator(); operator != nil && len(item.Name) > 0 {
				if value := item.getValue(); value != nil {
					where := fmt.Sprintf("%s %s ?", item.Name, *operator)

					query = query.Where(where, value)
				}
			}
		}
	}

	return query
}

func (a arguments) attachOR(query *gorm.DB) *gorm.DB {
	if a.Or != nil {
		for _, item := range *a.Or {
			if operator := item.resolveOperator(); operator != nil && len(item.Name) > 0 {
				if value := item.getValue(); value != nil {
					where := fmt.Sprintf("%s %s ?", item.Name, *operator)

					query = query.Or(where, value)
				}
			}
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

func (f *field) resolveOperator() *string {
	var operator string

	value := f.getValue()

	if f.Eq != nil {
		switch value.(type) {
		case string, *string:
			if isPostgres() {
				operator = "ILIKE"
			}

			operator = "ILIKE"
		default:
			operator = "="
		}
	} else if f.Ne != nil {
		switch value.(type) {
		case string, *string:
			if isPostgres() {
				operator = "NOT ILIKE"
			}

			operator = "NOT LIKE"
		default:
			operator = "<>"
		}
	} else if f.Gt != nil {
		operator = ">"
	} else if f.Lt != nil {
		operator = "<"
	} else if f.Gte != nil {
		operator = ">="
	} else if f.Lte != nil {
		operator = "<="
	}

	if len(operator) > 0 {
		return &operator
	}

	return nil
}

func toCamelCase(str string) string {
	camel := strcase.ToCamel(str)

	if strings.HasSuffix(camel, "Id") {
		camel = strings.TrimSuffix(camel, "Id")
		camel = camel + "ID"
	}

	return camel
}

func isValidType(kind reflect.Kind, valueType reflect.Type) bool {
	valueKind := valueType.Kind()
	valueElemKind := valueType.Elem().Kind()
	ptrKind := reflect.Ptr
	intKind := reflect.Int
	floatKind := reflect.Float64

	return (valueKind == ptrKind && valueElemKind == kind) ||
		(valueKind == ptrKind && valueElemKind == floatKind && kind == intKind) ||
		(valueKind == ptrKind && valueElemKind == intKind && kind == floatKind) ||
		(valueKind == kind) ||
		(valueKind == floatKind && kind == intKind) ||
		(valueKind == intKind && kind == floatKind)
}

func isValidNumericType(valueType reflect.Type) bool {
	valueKind := valueType.Kind()
	valueElemKind := valueType.Elem().Kind()
	intKind := reflect.Int
	floatKind := reflect.Float64

	return (valueKind == reflect.Ptr && (valueElemKind == intKind || valueElemKind == floatKind)) ||
		(valueKind == intKind || valueKind == floatKind)
}

func isNumericKind(kind reflect.Kind) bool {
	return (kind == reflect.Int) ||
		(kind == reflect.Int8) ||
		(kind == reflect.Int16) ||
		(kind == reflect.Int32) ||
		(kind == reflect.Int64) ||
		(kind == reflect.Uint) ||
		(kind == reflect.Float32) ||
		(kind == reflect.Float64)
}
