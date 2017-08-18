package controller

import (
	"hermes/database"
	"hermes/stats/parser"
	"hermes/stats/responses"

	"github.com/labstack/echo"
)

type frame struct {
	request *parser.Request
	context echo.Context
}

// PostStats is the main GraphQL controller
func PostStats(context echo.Context) error {
	request, err := parser.Parse(context)

	if err != nil {
		return err
	}

	dbs := database.GetReadDB()
	frame := &frame{request: request, context: context}

	defer db.Close()

	// Do something...

	return responses.PostResponse(frame.context)
}
