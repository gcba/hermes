package main

import (
	"fmt"
	"strconv"

	ratingsController "hermes/ratings/controller"
	ratingsHandler "hermes/ratings/handler"
	statsController "hermes/stats/controller"
	statsHandler "hermes/stats/handler"
	"hermes/utils"

	"github.com/alecthomas/kingpin"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo"
)

var (
	app            = kingpin.New("hermes", "GCBA product ratings APIs.")
	startCommand   = kingpin.Command("start", "Start an Hermes API.")
	ratingsCommand = startCommand.Command("ratings", "Start the ratings API.")
	statsCommand   = startCommand.Command("stats", "Start the statistics API.")
	hideCursor     = "\n\n\033[?25l"
	showCursor     = "\033[?25h"
	banner         = `
 _  _ ____ ____ _  _ ____ ____
 |__| |___ |__/ |\/| |___ [__
 |  | |___ |  \ |  | |___ ___] `
)

func main() {
	kingpin.Version("0.0.1")
	fmt.Print("\n", banner, "\n\n\n")

	switch kingpin.Parse() {
	case ratingsCommand.FullCommand():
		startRatingsAPI()
	case statsCommand.FullCommand():
		startStatsAPI()
	}
}

func startRatingsAPI() {
	routes := map[string]echo.HandlerFunc{
		"OptionsRoot":    ratingsController.OptionsRoot,
		"OptionsRatings": ratingsController.OptionsRatings,
		"PostRatings":    ratingsController.PostRatings}

	port := utils.GetPort("HERMES_RATINGS_PORT", 5000)
	handler := ratingsHandler.Handler(port, routes)

	fmt.Println("✅  Ratings server started on port", strconv.Itoa(port), hideCursor)
	gracehttp.Serve(handler.Server)
	fmt.Print(showCursor)
}

func startStatsAPI() {
	routes := map[string]echo.HandlerFunc{"PostStats": statsController.PostStats}
	port := utils.GetPort("HERMES_STATS_PORT", 7000)
	handler := statsHandler.Handler(port, routes)

	fmt.Println("✅  Stats server started on port", strconv.Itoa(port), hideCursor)
	gracehttp.Serve(handler.Server)
	fmt.Print(showCursor)
}
