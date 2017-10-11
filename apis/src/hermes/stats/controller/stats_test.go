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

	r := e.POST("/ratings").
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		WithHeader("Accept", "application/json").
		WithText(query).
		Expect()

	r.Status(http.StatusOK)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")

	json := r.JSON().Object()

	json.ContainsMap(meta)
	json.ContainsKey("data").ContainsKey("count")
	json.NotContainsKey("errors")
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

	r := e.POST("/ratings").
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		WithHeader("Accept", "application/json").
		WithText(query).
		Expect()

	r.Status(http.StatusOK)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")

	json := r.JSON().Object()

	json.ContainsMap(meta)
	json.ContainsKey("data").ContainsKey("count")
	json.NotContainsKey("errors")
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

	r := e.POST("/ratings").
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		WithHeader("Accept", "application/json").
		WithText(query).
		Expect()

	r.Status(http.StatusOK)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")

	json := r.JSON().Object()

	json.ContainsMap(meta)
	json.ContainsKey("data").ContainsKey("count")
	json.NotContainsKey("errors")
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

	r := e.POST("/ratings").
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		WithHeader("Accept", "application/json").
		WithText(query).
		Expect()

	r.Status(http.StatusOK)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")

	json := r.JSON().Object()

	json.ContainsMap(meta)
	json.ContainsKey("data").ContainsKey("count")
	json.NotContainsKey("errors")
}

func TestCountInvalidTableBadRequest(t *testing.T) {
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
			  "name": "example"
		  }
		}
	  }
	`

	meta := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Bad Request"}}

	r := e.POST("/ratings").
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		WithHeader("Accept", "application/json").
		WithText(query).
		Expect()

	r.Status(http.StatusOK)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")

	json := r.JSON().Object()

	json.ContainsMap(meta)
	json.NotContainsKey("data")
	json.ContainsKey("errors")
}

func TestCountInvalidFieldBadRequest(t *testing.T) {
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
			  "name": "ratings.example"
		  }
		}
	  }
	`

	meta := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Bad Request"}}

	r := e.POST("/ratings").
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		WithHeader("Accept", "application/json").
		WithText(query).
		Expect()

	r.Status(http.StatusOK)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")

	json := r.JSON().Object()

	json.ContainsMap(meta)
	json.NotContainsKey("data")
	json.ContainsKey("errors")
}

func TestAverage(t *testing.T) {
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
		"query": "query Example($field: Field!) { average(field: $field) }",
		"variables": {
		  "field": {
			  "name": "ratings.rating"
		  }
		}
	  }
	`

	meta := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    http.StatusOK,
			"message": "OK"}}

	r := e.POST("/ratings").
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		WithHeader("Accept", "application/json").
		WithText(query).
		Expect()

	r.Status(http.StatusOK)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")

	json := r.JSON().Object()

	json.ContainsMap(meta)
	json.ContainsKey("data").ContainsKey("average")
	json.NotContainsKey("errors")
}

func TestAverageAnd(t *testing.T) {
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
		"query": "query Example($field: Field!, $and: [Field!]) { average(field: $field, and: $and) }",
		"variables": {
		  "field": {
			  "name": "ratings.rating"
		  },
		  "and": {
			  "name": "ratings.has_message",
			  "eq": true
		  }
		}
	  }
	`

	meta := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    http.StatusOK,
			"message": "OK"}}

	r := e.POST("/ratings").
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		WithHeader("Accept", "application/json").
		WithText(query).
		Expect()

	r.Status(http.StatusOK)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")

	json := r.JSON().Object()

	json.ContainsMap(meta)
	json.ContainsKey("data").ContainsKey("average")
	json.NotContainsKey("errors")
}

func TestAverageOr(t *testing.T) {
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
		"query": "query Example($field: Field!, $or: [Field!]) { average(field: $field, or: $or) }",
		"variables": {
		  "field": {
			  "name": "ratings.rating"
		  },
		  "or": {
			"name": "ratings.has_message",
			"eq": true
		  }
		}
	  }
	`

	meta := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    http.StatusOK,
			"message": "OK"}}

	r := e.POST("/ratings").
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		WithHeader("Accept", "application/json").
		WithText(query).
		Expect()

	r.Status(http.StatusOK)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")

	json := r.JSON().Object()

	json.ContainsMap(meta)
	json.ContainsKey("data").ContainsKey("average")
	json.NotContainsKey("errors")
}

func TestAverageInvalidTableBadRequest(t *testing.T) {
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
		"query": "query Example($field: Field!) { average(field: $field) }",
		"variables": {
		  "field": {
			  "name": "example"
		  }
		}
	  }
	`

	meta := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Bad Request"}}

	r := e.POST("/ratings").
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		WithHeader("Accept", "application/json").
		WithText(query).
		Expect()

	r.Status(http.StatusOK)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")

	json := r.JSON().Object()

	json.ContainsMap(meta)
	json.NotContainsKey("data")
	json.ContainsKey("errors")
}

func TestAverageInvalidFieldBadRequest(t *testing.T) {
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
		"query": "query Example($field: Field!) { average(field: $field) }",
		"variables": {
		  "field": {
			  "name": "ratings.example"
		  }
		}
	  }
	`

	meta := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Bad Request"}}

	r := e.POST("/ratings").
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		WithHeader("Accept", "application/json").
		WithText(query).
		Expect()

	r.Status(http.StatusOK)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")

	json := r.JSON().Object()

	json.ContainsMap(meta)
	json.NotContainsKey("data")
	json.ContainsKey("errors")
}

func TestAverageNoFieldBadRequest(t *testing.T) {
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
		"query": "query Example($field: Field!) { average(field: $field) }",
		"variables": {
		  "field": {
			  "name": "ratings"
		  }
		}
	  }
	`

	meta := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Bad Request"}}

	r := e.POST("/ratings").
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		WithHeader("Accept", "application/json").
		WithText(query).
		Expect()

	r.Status(http.StatusOK)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")

	json := r.JSON().Object()

	json.ContainsMap(meta)
	json.NotContainsKey("data")
	json.ContainsKey("errors")
}

func getPort(env string, defaultPort int) string {
	port := os.Getenv(env)

	if len(port) == 0 {
		return string(defaultPort)
	}

	return port
}
