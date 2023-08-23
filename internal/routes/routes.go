package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lucasscarioca/music-stash/internal/controllers"
	"github.com/lucasscarioca/music-stash/internal/routes/middlewares"
	"github.com/lucasscarioca/music-stash/internal/views"
)

func MountRoutes(e *echo.Echo) {
	// Serve Static Files
	e.Use(middlewares.CacheControl(0), middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "assets/static",
		Browse: false,
	}))

	// Setup Templates
	e.Renderer = views.WithTemplateRegistry()

	e.GET("ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	// Pages Routes
	pages := e.Group("/", middlewares.CacheControl(0))
	pages.GET("", controllers.Home)
	pages.GET("404", func(c echo.Context) error {
		return c.Render(http.StatusOK, "pageNotFound", map[string]any{})
	})

	// Htmx Fragments Routes
	hx := e.Group("/hx", middlewares.CacheControl(0), middlewares.ValidateHxRequest)
	hx.GET("/time", controllers.Time)
}
