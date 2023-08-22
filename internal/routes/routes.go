package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lucasscarioca/music-stash/internal/controllers"
	"github.com/lucasscarioca/music-stash/internal/routes/middlewares"
)

func MountRoutes(e *echo.Echo) {
	// Serve Static Files
	e.Use(middlewares.CacheControl(0), middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "assets/static",
		Browse: false,
	}))

	// Setup Templates
	e.Renderer = &templateRegistry{
		templates: registerTemplates(),
	}

	e.GET("ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	// Pages Routes
	pages := e.Group("/", middlewares.CacheControl(0))
	pages.GET("", controllers.Home)
	pages.GET("404", func(c echo.Context) error {
		return c.Render(http.StatusOK, "404.html", map[string]any{})
	})

	// Components Routes
	components := e.Group("/components", middlewares.CacheControl(0))
	components.GET("/time", controllers.Time)

}
