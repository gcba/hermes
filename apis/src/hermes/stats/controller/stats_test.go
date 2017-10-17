package controller

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"hermes/stats/handler"
	"hermes/utils"

	"github.com/gavv/httpexpect"
	"github.com/labstack/echo"
)

var (
	routes = map[string]echo.HandlerFunc{"PostStats": PostStats}
	port   = strconv.Itoa(utils.GetPort("HERMES_STATS_PORT", 7000))
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

	r := e.POST("/stats").
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		WithHeader("Accept", "application/json").
		WithText(query).
		Expect()

	r.Status(http.StatusOK)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")

	json := r.JSON().Object()

	json.ContainsMap(meta)
	json.ContainsKey("data")
	json.NotContainsKey("errors")
}

func TestCount_WithOperator(t *testing.T) {
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

	r := e.POST("/stats").
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		WithHeader("Accept", "application/json").
		WithText(query).
		Expect()

	r.Status(http.StatusOK)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")

	json := r.JSON().Object()

	json.ContainsMap(meta)
	json.ContainsKey("data")
	json.NotContainsKey("errors")
}

func TestCount_And(t *testing.T) {
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

	r := e.POST("/stats").
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		WithHeader("Accept", "application/json").
		WithText(query).
		Expect()

	r.Status(http.StatusOK)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")

	json := r.JSON().Object()

	json.ContainsMap(meta)
	json.ContainsKey("data")
	json.NotContainsKey("errors")
}

func TestCount_Or(t *testing.T) {
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

	r := e.POST("/stats").
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		WithHeader("Accept", "application/json").
		WithText(query).
		Expect()

	r.Status(http.StatusOK)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")

	json := r.JSON().Object()

	json.ContainsMap(meta)
	json.ContainsKey("data")
	json.NotContainsKey("errors")
}

func TestCount_InvalidTable_BadRequest(t *testing.T) {
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

	r := e.POST("/stats").
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		WithHeader("Accept", "application/json").
		WithText(query).
		Expect()

	r.Status(http.StatusBadRequest)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")

	json := r.JSON().Object()

	json.ContainsMap(meta)
	json.NotContainsKey("data")
	json.ContainsKey("errors")
}

func TestCount_InvalidField_BadRequest(t *testing.T) {
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

	r := e.POST("/stats").
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		WithHeader("Accept", "application/json").
		WithText(query).
		Expect()

	r.Status(http.StatusBadRequest)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")

	json := r.JSON().Object()

	json.ContainsMap(meta)
	json.NotContainsKey("data")
	json.ContainsKey("errors")
}

func TestCount_WrongFieldType_BadRequest(t *testing.T) {
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
				"name": "ratings.rating",
				"eq": true
		    }
		}
	}
	`

	meta := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Bad Request"}}

	r := e.POST("/stats").
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		WithHeader("Accept", "application/json").
		WithText(query).
		Expect()

	r.Status(http.StatusBadRequest)
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

	r := e.POST("/stats").
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		WithHeader("Accept", "application/json").
		WithText(query).
		Expect()

	r.Status(http.StatusOK)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")

	json := r.JSON().Object()

	json.ContainsMap(meta)
	json.ContainsKey("data")
	json.NotContainsKey("errors")
}

func TestAverage_WithOperator(t *testing.T) {
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
				"name": "ratings.rating",
				"gt": 0
		    }
		}
	}
	`

	meta := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    http.StatusOK,
			"message": "OK"}}

	r := e.POST("/stats").
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		WithHeader("Accept", "application/json").
		WithText(query).
		Expect()

	r.Status(http.StatusOK)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")

	json := r.JSON().Object()

	json.ContainsMap(meta)
	json.ContainsKey("data")
	json.NotContainsKey("errors")
}

func TestAverage_And(t *testing.T) {
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
				"name": "ratings.rating",
				"gt": 0
		    },
		    "and": {
			    "name": "ratings.rating",
			    "lt": 4
		    }
		}
	}
	`

	meta := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    http.StatusOK,
			"message": "OK"}}

	r := e.POST("/stats").
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		WithHeader("Accept", "application/json").
		WithText(query).
		Expect()

	r.Status(http.StatusOK)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")

	json := r.JSON().Object()

	json.ContainsMap(meta)
	json.ContainsKey("data")
	json.NotContainsKey("errors")
}

func TestAverage_Or(t *testing.T) {
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
				"name": "ratings.rating",
				"gt": 3
		  	},
		  	"or": {
				"name": "ratings.rating",
				"lte": 2
			}
		}
	}
	`

	meta := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    http.StatusOK,
			"message": "OK"}}

	r := e.POST("/stats").
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		WithHeader("Accept", "application/json").
		WithText(query).
		Expect()

	r.Status(http.StatusOK)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")

	json := r.JSON().Object()

	json.ContainsMap(meta)
	json.ContainsKey("data")
	json.NotContainsKey("errors")
}

func TestAverage_InvalidTable_BadRequest(t *testing.T) {
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

	r := e.POST("/stats").
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		WithHeader("Accept", "application/json").
		WithText(query).
		Expect()

	r.Status(http.StatusBadRequest)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")

	json := r.JSON().Object()

	json.ContainsMap(meta)
	json.NotContainsKey("data")
	json.ContainsKey("errors")
}

func TestAverage_InvalidField_BadRequest(t *testing.T) {
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

	r := e.POST("/stats").
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		WithHeader("Accept", "application/json").
		WithText(query).
		Expect()

	r.Status(http.StatusBadRequest)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")

	json := r.JSON().Object()

	json.ContainsMap(meta)
	json.NotContainsKey("data")
	json.ContainsKey("errors")
}

func TestAverage_NoField_BadRequest(t *testing.T) {
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
			    "name": "stats"
		    }
		}
	}
	`

	meta := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Bad Request"}}

	r := e.POST("/stats").
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		WithHeader("Accept", "application/json").
		WithText(query).
		Expect()

	r.Status(http.StatusBadRequest)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")

	json := r.JSON().Object()

	json.ContainsMap(meta)
	json.NotContainsKey("data")
	json.ContainsKey("errors")
}

func TestAverage_NonNumericField_BadRequest(t *testing.T) {
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
			    "name": "stats.has_message"
		    }
		}
	}
	`

	meta := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Bad Request"}}

	r := e.POST("/stats").
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		WithHeader("Accept", "application/json").
		WithText(query).
		Expect()

	r.Status(http.StatusBadRequest)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")

	json := r.JSON().Object()

	json.ContainsMap(meta)
	json.NotContainsKey("data")
	json.ContainsKey("errors")
}
