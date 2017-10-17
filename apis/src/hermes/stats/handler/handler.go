package handler

import (
	"crypto/rsa"
	"io/ioutil"
	"os"
	"strconv"

	"hermes/middlewares"
	"hermes/responses"
	"hermes/stats/schema"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/ssh-vault/ssh2pem"
)

type RequestValidator struct {
	validator *validator.Validate
}

func (rv *RequestValidator) Validate(request interface{}) error {
	return rv.validator.Struct(request)
}

func Handler(port int, routes map[string]echo.HandlerFunc) *echo.Echo {
	schema.Parse()

	e := echo.New()
	validate := validator.New()
	env := os.Getenv("HERMES_STATS_ENV")
	key := getKey(e)

	jwtConfig := middleware.JWTConfig{
		SigningKey:    key,
		SigningMethod: "RS256",
		ContextKey:    "jwt"}

	if env == "DEV" {
		e.Logger.SetLevel(log.DEBUG)

		e.Debug = true
	} else {
		e.Logger.SetLevel(log.ERROR)
		e.Pre(middleware.HTTPSRedirect())
		e.Use(middleware.JWTWithConfig(jwtConfig))
	}

	e.Use(middleware.Secure())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("20K"))
	e.Use(middlewares.NotImplementedMiddleware)
	e.Use(middlewares.NotAcceptableMiddleware)
	e.Use(middlewares.BadRequestMiddleware)
	e.Use(middlewares.UnsupportedMediaTypeMiddleware)
	e.Use(middlewares.CorsMiddleware)

	e.POST("/stats", routes["PostStats"])

	e.HTTPErrorHandler = responses.ErrorHandler
	e.Validator = &RequestValidator{validator: validate}
	e.Server.Addr = ":" + strconv.Itoa(port)

	return e
}

func getKey(echo *echo.Echo) *rsa.PublicKey { // TODO: Move to a shared utils package
	keyArray, readErr := ioutil.ReadFile(os.Getenv("HERMES_STATS_KEY"))

	if readErr != nil {
		echo.Logger.Fatal("Could not find key file")
	}

	keyPEM, sshParseErr := ssh2pem.GetPublicKeyPem(string(keyArray[:]))

	if sshParseErr != nil {
		echo.Logger.Fatal("Could not parse SSH key")
	}

	key, parseErr := jwt.ParseRSAPublicKeyFromPEM(keyPEM)

	if parseErr != nil {

		echo.Logger.Fatal("Invalid PEM key")
	}

	/*
		token := jwt.New(jwt.SigningMethodRS256)
		privKey, _ := ioutil.ReadFile(os.Getenv("HERMES_STATS_PRIVATEKEY"))
		privKeyParsed, _ := jwt.ParseRSAPrivateKeyFromPEM(privKey)

		t, _ := token.SignedString(privKeyParsed)

		echo.Logger.Fatal(t)
	*/

	return key
}
