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

	json := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    200,
			"message": "OK"},
		"endpoints": []map[string]interface{}{
			{
				"method": "POST",
				"path":   "/ratings",
				"headers": [][]string{
					{
						"Content-Type",
						"application/json; charset=utf-8"}}}}}

	e.OPTIONS("/ratings").WithJSON(json).
		WithHeader("Content-Type", "application/json; charset=utf-8").
		WithHeader("Allow", "OPTIONS POST").
		Expect().
		Status(http.StatusOK)
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

	jsonRequest := map[string]interface{}{
		"rating":      uint8(3),
		"description": "Regular",
		"range":       "e10adc3949ba59abbe56e057f20f883e",
		"app": map[string]interface{}{
			"key":     "e10adc3949ba59abbe56e057f20f883e",
			"version": "2.O"},
		"platform": map[string]interface{}{
			"key":     "e10adc3949ba59abbe56e057f20f883e",
			"version": "9.O"}}

	/*
		jsonResponse := map[string]interface{}{
			"meta": map[string]interface{}{
				"code":    201,
				"message": "Created"}}}
	*/

	e.POST("/ratings").WithJSON(jsonRequest).
		WithHeader("Content-Type", "application/json; charset=utf-8").
		Expect().
		Status(http.StatusCreated)
}