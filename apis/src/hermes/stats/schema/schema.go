package schema

import (
	"errors"
	"io/ioutil"
	"path"
	"runtime"

	"github.com/neelance/graphql-go"
)

const DB = iota

var Schema *graphql.Schema

type Value struct {
	String string
	Int    int
	Float  float64
	Bool   bool
}

func (_ Value) ImplementsGraphQLType(name string) bool {
	return name == "Value"
}

func (v *Value) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		v.String = input
	case int:
		v.Int = input
	case float64:
		v.Float = input
	case bool:
		v.Bool = input
	default:
		return errors.New("Wrong type")
	}

	return nil
}

func Parse() {
	var rawSchema []byte
	var err error

	_, filename, _, ok := runtime.Caller(1)

	if !ok {
		panic("Could not load GraphQL schema")
	}

	rawSchema, err = ioutil.ReadFile(path.Join(path.Dir(filename), "../schema/schema.graphql"))

	if err != nil {
		panic(err)
	}

	Schema, err = graphql.ParseSchema(string(rawSchema), &Resolver{})

	if err != nil {
		panic(err)
	}
}
