package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/lucasscarioca/music-stash/configs"
	"github.com/lucasscarioca/music-stash/internal/routes"
	"github.com/lucasscarioca/music-stash/internal/routes/middlewares"
)

func main() {
	configs.Load()

	// db.Connect()

	e := echo.New()
	middlewares.Mount(e)
	routes.Mount(e)

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
