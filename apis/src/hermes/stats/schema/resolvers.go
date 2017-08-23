package schema

import (
	"context"
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
		Name     string
		Operator *string
		Value    Value
		Next     *operation
	}

	operation struct {
		Condition string
		Field     field
	}

	Resolver struct{}
)

var modelList = map[string]interface{}{
	"apps":      []models.App{},
	"appusers":  []models.AppUser{},
	"brands":    []models.Brand{},
	"browsers":  []models.Browser{},
	"devices":   []models.Device{},
	"messages":  []models.Message{},
	"platforms": []models.Platform{},
	"ranges":    []models.Range{},
	"ratings":   []models.Rating{},
}

/*
var modelList = map[string]reflect.Type{
	"apps":      reflect.TypeOf(models.App{}),
	"appusers":  reflect.TypeOf(models.AppUser{}),
	"brands":    reflect.TypeOf(models.Brand{}),
	"browsers":  reflect.TypeOf(models.Browser{}),
	"devices":   reflect.TypeOf(models.Device{}),
	"messages":  reflect.TypeOf(models.Message{}),
	"platforms": reflect.TypeOf(models.Platform{}),
	"ranges":    reflect.TypeOf(models.Range{}),
	"ratings":   reflect.TypeOf(models.Rating{}),
}
*/

func errorResponse() error {
	return echo.NewHTTPError(http.StatusInternalServerError)
}

func (r *Resolver) Count(context context.Context, args struct{ Field field }) (int32, error) {
	// entity := getEntity(args.Field.Name)

	return 0, nil
}

func (r *Resolver) Average(context context.Context, args struct{ Field field }) (float64, error) {
	// TODO: Implement

	return 0, nil
}

func (f *field) flatten(buffer []*field) []*field {
	if f.Next != nil {
		f.Next.Field.flatten(buffer)
	}

	f.Next = nil
	buffer = append(buffer, f)

	return buffer
}

func (f *field) query(context context.Context) *gorm.DB {
	if db, castOk := context.Value(DB).(*gorm.DB); castOk {
		if model, modelExists := f.resolveModel(); modelExists {
			entity := f.getEntity()

			return db.Where(map[string]interface{}{entity.Field: f.Value}).Find(&model)
		}
	}

	// TODO: Handle non existent db

	return nil
}

func (f *field) getEntity() entity {
	splitField := strings.Split(f.Name, ".")

	return entity{Table: splitField[0], Field: splitField[1]}
}

func (f *field) resolveOperator() (string, bool) {
	if f.Operator != nil {
		var result string

		switch *f.Operator {
		case "EQ":
			result = "="

			return result, true
		}
	}

	return "", false
}

func (f *field) resolveModel() (interface{}, bool) {
	entity := f.getEntity()

	if model, ok := modelList[entity.Table]; ok {
		return model, true
	}

	// TODO: Handle non existent model

	return nil, false
}

/*
func (f *field) resolveModel() (interface{}, bool) {
	entity := f.getEntity()

	if model, ok := modelList[entity.Table]; ok {
		result := reflect.New(model).Elem().Interface()

		return result, true
	}

	// TODO: Handle non existent model

	return nil, false
}
*/
