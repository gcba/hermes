package main

import (
	"fmt"
	"os"
	"strconv"

	ratingsController "hermes/ratings/controller"
	ratingsHandler "hermes/ratings/handler"
	statsController "hermes/stats/controller"
	statsHandler "hermes/stats/handler"

	"github.com/alecthomas/kingpin"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo"
)

var (
	app            = kingpin.New("hermes", "GCBA product ratings APIs.")
	startCommand   = kingpin.Command("start", "Start an Hermes API.")
	ratingsCommand = startCommand.Command("ratings", "Start the ratings API.")
	statsCommand   = startCommand.Command("stats", "Start the statistics API.")
	ratingsPort    = getPort("HERMES_RATINGS_PORT", 5000)
	statsPort      = getPort("HERMES_STATS_PORT", 7000)
	noCursor       = "\n\n\033[?25l"
	banner         = `
 _  _ ____ ____ _  _ ____ ____
 |__| |___ |__/ |\/| |___ [__
 |  | |___ |  \ |  | |___ ___] `
)

func main() {
	kingpin.Version("0.0.1")
	fmt.Println("\n", banner)

	switch kingpin.Parse() {
	case "start ratings":
		fmt.Print("	               ratings", "\n\n\n")
		startRatingsAPI()
	case "start stats":
		fmt.Print("	                 stats", "\n\n\n")
		startStatsAPI()
	}
}

func startRatingsAPI() {
	routes := map[string]echo.HandlerFunc{
		"OptionsRoot":    ratingsController.OptionsRoot,
		"OptionsRatings": ratingsController.OptionsRatings,
		"PostRatings":    ratingsController.PostRatings}

	handler, castOk := ratingsHandler.Handler(ratingsPort, routes).(*echo.Echo)

	if !castOk {
		handler.Logger.Fatal("Could not start server")
	}

	fmt.Println("✅  Ratings server started on port", strconv.Itoa(ratingsPort))
	fmt.Print(noCursor)

	handler.Logger.Fatal(gracehttp.Serve(handler.Server))
}

func startStatsAPI() {
	routes := map[string]echo.HandlerFunc{"PostStats": statsController.PostStats}
	handler, castOk := statsHandler.Handler(statsPort, routes).(*echo.Echo)

	if !castOk {
		handler.Logger.Fatal("Could not start server")
	}

	fmt.Println("✅  Stats server started on port", strconv.Itoa(statsPort))
	fmt.Print(noCursor)

	handler.Logger.Fatal(gracehttp.Serve(handler.Server))
}

func getPort(env string, defaultPort int) int {
	port, portErr := strconv.Atoi(os.Getenv(env))

	if portErr != nil {
		return defaultPort
	}

	return port
}
