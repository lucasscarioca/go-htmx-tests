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

	root := e.Group("/", middlewares.CacheControl(0))

	root.GET("ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	// Templates
	root.GET("", controllers.Home)

	// Components

	root.GET("404", func(c echo.Context) error {
		return c.Render(http.StatusOK, "404.html", map[string]any{})
	})
}
