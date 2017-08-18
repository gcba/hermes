package controller

import (
	"hermes/database"
	"hermes/stats/parser"
	"hermes/stats/responses"

	"github.com/labstack/echo"
)

// PostStats is the main GraphQL controller
func PostStats(context echo.Context) error {
	request, err := parser.Parse(context)

	if err != nil {
		return err
	}

	dbs := &databases{read: database.GetWriteDB(), write: database.GetReadDB()}
	frame := &frame{request: request, context: context}

	defer dbs.read.Close()
	defer dbs.write.Close()

	// Do something...

	return responses.PostResponse(frame.context)
}
