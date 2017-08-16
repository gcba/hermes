package main

import (
	"fmt"
	"os"
	"strconv"

	"hermes/ratings/controller"
	"hermes/ratings/handler"

	"github.com/alecthomas/kingpin"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo"
)

var (
	app            = kingpin.New("hermes", "GCBA product ratings APIs.")
	startCommand   = kingpin.Command("start", "Start an Hermes API.")
	ratingsCommand = startCommand.Command("ratings", "Start the ratings API.")
	statsCommand   = startCommand.Command("stats", "Start the statistics API.")
	ratingsPort    = getRatingsPort()
	noCursor       = "\n\n\033[?25l"
	banner         = `
 _  _ ____ ____ _  _ ____ ____
 |__| |___ |__/ |\/| |___ [__
 |  | |___ |  \ |  | |___ ___] `
)

func main() {
	kingpin.Version("0.0.1")
	fmt.Println("\n", banner, "\n\n")

	switch kingpin.Parse() {
	case "start ratings":
		startRatingsAPI()
	case "start stats":
		startStatsAPI()
	}
}

func startRatingsAPI() {
	routes := map[string]echo.HandlerFunc{
		"OptionsRoot":    controller.OptionsRoot,
		"OptionsRatings": controller.OptionsRatings,
		"PostRatings":    controller.PostRatings}

	handler, castOk := handler.Handler(ratingsPort, routes).(*echo.Echo)

	if !castOk {
		handler.Logger.Fatal("Could not start server")
	}

	fmt.Println("✅  Server started on port", strconv.Itoa(ratingsPort))
	fmt.Print(noCursor)

	handler.Logger.Fatal(gracehttp.Serve(handler.Server))
}

func startStatsAPI() {
	routes := map[string]echo.HandlerFunc{
		"OptionsRoot":    controller.OptionsRoot,
		"OptionsRatings": controller.OptionsRatings,
		"PostRatings":    controller.PostRatings}

	handler, castOk := handler.Handler(ratingsPort, routes).(*echo.Echo)

	if !castOk {
		handler.Logger.Fatal("Could not start server")
	}

	fmt.Println("✅  Server started on port", strconv.Itoa(ratingsPort))
	fmt.Print(noCursor)

	handler.Logger.Fatal(gracehttp.Serve(handler.Server))
}

func getRatingsPort() int {
	port, portErr := strconv.Atoi(os.Getenv("HERMES_RATINGS_PORT"))

	if portErr != nil {
		return 5000
	}

	return port
}

func getStatsPort() int {
	port, portErr := strconv.Atoi(os.Getenv("HERMES_STATS_PORT"))

	if portErr != nil {
		return 5000
	}

	return port
}
