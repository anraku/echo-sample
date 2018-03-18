package main

import (
	auth "github.com/anraku/echo-sample/auth"
	handler "github.com/anraku/echo-sample/handler"
	log "github.com/anraku/echo-sample/log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Pre(middleware.Logger())
	e.Pre(middleware.LoggerWithConfig(log.OutputAppLog("/var/log/echo-sample/api.log")))
	e.Pre(auth.AuthToken())
	e.Pre(auth.ServerHeader())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/json", handler.SendJSON)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
