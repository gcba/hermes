package handler

import (
	"os"
	"strconv"

	"hermes/responses"
	"hermes/utils"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func Handler(port int, env string, publicKey string) *echo.Echo {
	e := echo.New()

	if os.Getenv(env) == "DEV" {
		e.Debug = true

		e.Logger.SetLevel(log.DEBUG)
	} else {
		jwtConfig := middleware.JWTConfig{
			SigningKey:    utils.GetPublicKey(publicKey, e),
			SigningMethod: "RS256",
			ContextKey:    "jwt"}

		e.Logger.SetLevel(log.ERROR)
		e.Pre(middleware.HTTPSRedirect())
		e.Use(middleware.JWTWithConfig(jwtConfig))
	}

	e.Use(middleware.Secure())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("20K"))
	e.Use(NotImplementedMiddleware)
	e.Use(NotAcceptableMiddleware)
	e.Use(BadRequestMiddleware)
	e.Use(UnsupportedMediaTypeMiddleware)
	e.Use(CorsMiddleware)

	e.HTTPErrorHandler = responses.ErrorHandler
	e.Server.Addr = ":" + strconv.Itoa(port)

	return e
}
