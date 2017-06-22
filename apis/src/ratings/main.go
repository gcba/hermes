package main

import (
	"fmt"
	"strconv"

	"ratings/controller"
	"ratings/handler"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo"
)

func main() {
	port := 3000

	routes := map[string]echo.HandlerFunc{
		"OptionsRoot":    controller.OptionsRoot,
		"OptionsRatings": controller.OptionsRatings,
		"PostRatings":    controller.PostRatings}

	handler := handler.Handler(port, routes).(*echo.Echo) // Casting via type assertion

	fmt.Println("Started server on port", strconv.Itoa(port))

	handler.Logger.Fatal(gracehttp.Serve(handler.Server))
}
