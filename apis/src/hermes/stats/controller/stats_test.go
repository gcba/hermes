package controller

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"hermes/ratings/handler"

	"github.com/gavv/httpexpect"
	"github.com/labstack/echo"
)

var (
	routes = map[string]echo.HandlerFunc{"PostStats": PostStats}
	port   = getPort("HERMES_STATS_PORT", 7000)
)

func TestCount(t *testing.T) {
	handler := handler.Handler(3000, routes)
	server := httptest.NewServer(handler)

	defer server.Close()

	server.URL = "http://localhost:" + port

	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  server.URL,
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	query := `
	{
		"query": "query Example($field: Field!) { count(field: $field) }",
		"variables": {
		  "field": {
			  "name": "messages"
		  }
		}
	  }
	`

	meta := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    http.StatusOK,
			"message": "OK"}}

	data := map[string]interface{}{}

	r := e.POST("/ratings").
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		WithHeader("Accept", "application/json").
		WithText(query).
		Expect()

	r.Status(http.StatusOK)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")

	json := r.JSON().Object()

	json.ContainsMap(meta)
	json.ContainsMap(data).ContainsKey("count")
}

func TestCountWithOperator(t *testing.T) {
	handler := handler.Handler(3000, routes)
	server := httptest.NewServer(handler)

	defer server.Close()

	server.URL = "http://localhost:" + port

	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  server.URL,
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	query := `
	{
		"query": "query Example($field: Field!) { count(field: $field) }",
		"variables": {
		  "field": {
			  "name": "messages.status",
			  "eq": 0
		  }
		}
	  }
	`

	meta := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    http.StatusOK,
			"message": "OK"}}

	data := map[string]interface{}{}

	r := e.POST("/ratings").
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		WithHeader("Accept", "application/json").
		WithText(query).
		Expect()

	r.Status(http.StatusOK)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")

	json := r.JSON().Object()

	json.ContainsMap(meta)
	json.ContainsMap(data).ContainsKey("count")
}

func TestCountAnd(t *testing.T) {
	handler := handler.Handler(3000, routes)
	server := httptest.NewServer(handler)

	defer server.Close()

	server.URL = "http://localhost:" + port

	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  server.URL,
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	query := `
	{
		"query": "query Example($field: Field!, $and: [Field!]) { count(field: $field, and: $and) }",
		"variables": {
		  "field": {
			  "name": "messages.status",
			  "eq": 0
		  },
		  "and": {
			  "name": "messages.rating_id",
			  "lte": 5
		  }
		}
	  }
	`

	meta := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    http.StatusOK,
			"message": "OK"}}

	data := map[string]interface{}{}

	r := e.POST("/ratings").
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		WithHeader("Accept", "application/json").
		WithText(query).
		Expect()

	r.Status(http.StatusOK)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")

	json := r.JSON().Object()

	json.ContainsMap(meta)
	json.ContainsMap(data).ContainsKey("count")
}

func TestCountOr(t *testing.T) {
	handler := handler.Handler(3000, routes)
	server := httptest.NewServer(handler)

	defer server.Close()

	server.URL = "http://localhost:" + port

	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  server.URL,
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	query := `
	{
		"query": "query Example($field: Field!, $or: [Field!]) { count(field: $field, or: $or) }",
		"variables": {
		  "field": {
			  "name": "messages.status",
			  "eq": 0
		  },
		  "or": {
			  "name": "messages.rating_id",
			  "lte": 5
		  }
		}
	  }
	`

	meta := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    http.StatusOK,
			"message": "OK"}}

	data := map[string]interface{}{}

	r := e.POST("/ratings").
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		WithHeader("Accept", "application/json").
		WithText(query).
		Expect()

	r.Status(http.StatusOK)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")

	json := r.JSON().Object()

	json.ContainsMap(meta)
	json.ContainsMap(data).ContainsKey("count")
}

func getPort(env string, defaultPort int) string {
	port := os.Getenv(env)

	if len(port) == 0 {
		return string(defaultPort)
	}

	return port
}
