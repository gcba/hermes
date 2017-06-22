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
						"application/json; charset=utf-8"},
					{
						"Allow",
						"OPTIONS POST"}}},
			{
				"method": "POST",
				"path":   "/ratings",
				"headers": [][]string{
					{
						"Content-Type",
						"application/json; charset=utf-8"}}}}}

	e.OPTIONS("/").WithJSON(json).
		WithHeader("Content-Type", "application/json; charset=utf-8").
		WithHeader("Allow", "OPTIONS").
		Expect().
		Status(http.StatusOK)
}
