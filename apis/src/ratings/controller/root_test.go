package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"ratings/handler"

	"github.com/gavv/httpexpect"
	"github.com/labstack/echo"
)

func TestOptionsRoot(t *testing.T) {
	routes := map[string]echo.HandlerFunc{
		"OptionsRoot":    OptionsRoot,
		"OptionsRatings": OptionsRatings,
		"PostRatings":    PostRatings}

	handler := handler.Handler(3000, routes)
	server := httptest.NewServer(handler)

	defer server.Close()

	server.URL = "http://localhost:3000"

	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  server.URL,
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	json := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    200,
			"message": "OK"},
		"endpoints": []map[string]interface{}{
			{
				"method": "OPTIONS",
				"path":   "/ratings",
				"headers": [][]string{
					{
						"Content-Type",
						"application/json; charset=UTF-8"},
					{
						"Allow",
						"OPTIONS POST"}}},
			{
				"method": "POST",
				"path":   "/ratings",
				"headers": [][]string{
					{
						"Content-Type",
						"application/json; charset=UTF-8"}}}}}

	e.OPTIONS("/").WithJSON(json).
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		WithHeader("Accept", "application/json").
		WithHeader("Allow", "OPTIONS").
		Expect().
		Status(http.StatusOK)
}

func TestOptionsRoot_BadRequestError(t *testing.T) {
	handler := handler.Handler(3000, routes)
	server := httptest.NewServer(handler)

	defer server.Close()

	server.URL = "http://localhost:3000"

	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  server.URL,
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	response := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Bad Request"},
		"errors": []interface{}{"Accept header is missing"}}

	r := e.OPTIONS("/").
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		Expect()

	r.Status(http.StatusBadRequest)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")
	r.JSON().Object().Equal(response)
}

func TestOptions_NotFoundError(t *testing.T) {
	handler := handler.Handler(3000, routes)
	server := httptest.NewServer(handler)

	defer server.Close()

	server.URL = "http://localhost:3000"

	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  server.URL,
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	response := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Not Found"}}

	r := e.OPTIONS("/example").
		WithHeader("Accept", "application/json").
		Expect()

	r.Status(http.StatusNotFound)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")
	r.JSON().Object().Equal(response)
}

func TestOptions_NotImplementedError(t *testing.T) {
	handler := handler.Handler(3000, routes)
	server := httptest.NewServer(handler)

	defer server.Close()

	server.URL = "http://localhost:3000"

	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  server.URL,
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	response := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    http.StatusNotImplemented,
			"message": "Not Implemented"}}

	r := e.GET("/").
		WithHeader("Accept", "application/json").
		Expect()

	r.Status(http.StatusNotImplemented)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")
	r.JSON().Object().Equal(response)
}
