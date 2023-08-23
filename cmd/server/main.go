package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lucasscarioca/music-stash/configs"
	"github.com/lucasscarioca/music-stash/internal/routes"
)

func main() {
	configs.Load()

	// db.Connect()

	e := echo.New()

	// Middlewares
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano} - ${uri} [${method} - ${status}] ${latency_human} - ${error}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.HTTPErrorHandler = customHTTPErrorHandler
	routes.MountRoutes(e)

	fmt.Printf("🚀 Starting server on port: %s\n", configs.GetPort())
	go func() {
		if err := e.Start(fmt.Sprintf(":%s", configs.GetPort())); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	host := c.Request().Host
	URI := c.Request().RequestURI
	qs := c.QueryString()

	c.Logger().Error(err, fmt.Sprintf("\non: %s%s%s error code: %d", host, URI, qs, code))
	if code == 404 {
		c.Redirect(http.StatusTemporaryRedirect, "/404")
	}
	c.String(code, fmt.Sprintf("error code: %d", code))
}
