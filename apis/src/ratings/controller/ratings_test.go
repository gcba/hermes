package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"ratings/handler"

	"github.com/gavv/httpexpect"
	"github.com/labstack/echo"
)

var routes = map[string]echo.HandlerFunc{
	"OptionsRoot":    OptionsRoot,
	"OptionsRatings": OptionsRatings,
	"PostRatings":    PostRatings}

func TestOptionsRatings(t *testing.T) {
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
			"code":    200,
			"message": "OK"},
		"endpoints": []map[string]interface{}{
			{
				"method": "POST",
				"path":   "/ratings",
				"headers": map[string]interface{}{
					"Content-Type": "application/json; charset=UTF-8"}}}}

	r := e.OPTIONS("/ratings").
		WithHeader("Accept", "application/json").
		Expect()

	r.Status(http.StatusOK)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")
	r.Header("Allow").Equal("OPTIONS POST")
	r.JSON().Object().Equal(response)
}

func TestPostRatings(t *testing.T) {
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

	request := map[string]interface{}{
		"rating":      uint8(3),
		"description": "Regular",
		"range":       "e10adc3949ba59abbe56e057f20f883e",
		"app": map[string]interface{}{
			"key":     "e10adc3949ba59abbe56e057f20f883e",
			"version": "2.O"},
		"platform": map[string]interface{}{
			"key":     "e10adc3949ba59abbe56e057f20f883e",
			"version": "9.O"}}

	response := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    201,
			"message": "Created"}}

	r := e.POST("/ratings").
		WithHeader("Content-Type", "application/json; charset=UTF-8").
		WithHeader("Accept", "application/json").
		WithJSON(request).
		Expect()

	r.Status(http.StatusCreated)
	r.Header("Content-Type").Equal("application/json; charset=UTF-8")
	r.JSON().Object().Equal(response)
}
