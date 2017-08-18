package schema

import (
	"io/ioutil"
	"path"
	"runtime"

	"hermes/database"

	"github.com/neelance/graphql-go"
)

const DB = iota

var Schema *graphql.Schema

func Parse() {
	var rawSchema []byte
	var err error

	db := database.GetReadDB()

	defer db.Close()

	_, filename, _, ok := runtime.Caller(1)

	if !ok {
		panic("Could not load GraphQL schema")
	}

	rawSchema, err = ioutil.ReadFile(path.Join(path.Dir(filename), "../schema/schema.graphql"))

	if err != nil {
		panic(err)
	}

	Schema, err = graphql.ParseSchema(string(rawSchema), NewResolver(db))

	if err != nil {
		panic(err)
	}
}
