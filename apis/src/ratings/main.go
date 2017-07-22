package main

import (
	"fmt"
	"os"
	"strconv"

	"ratings/controller"
	"ratings/handler"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo"
)

func main() {
	port, parseErr := strconv.Atoi(os.Getenv("API_RATINGS_PORT"))

	if parseErr != nil {
		port = 5000
	}

	routes := map[string]echo.HandlerFunc{
		"OptionsRoot":    controller.OptionsRoot,
		"OptionsRatings": controller.OptionsRatings,
		"PostRatings":    controller.PostRatings}

	handler, castOk := handler.Handler(port, routes).(*echo.Echo)

	if !castOk {
		handler.Logger.Fatal("Could not start server")
	}

	banner := `
_  _ ____ ____ _  _ ____ ____
|__| |___ |__/ |\/| |___ [__
|  | |___ |  \ |  | |___ ___] `

	fmt.Println("\n", banner)
	fmt.Println("\n\n✅  Server started on port", strconv.Itoa(port))
	fmt.Print("\n\n\033[?25l")

	handler.Logger.Fatal(gracehttp.Serve(handler.Server))
}
