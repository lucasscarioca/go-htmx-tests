package controllers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Todo struct {
	Title  string
	Status bool
}

func Home(c echo.Context) error {
	todos := map[string][]Todo{
		"Todos": {
			{Title: "brush teeth", Status: true},
			{Title: "Study", Status: false},
			{Title: "Dishes", Status: false},
		},
	}
	return c.Render(http.StatusOK, "home", todos)
}

func Time(c echo.Context) error {
	ctx := map[string]any{"ts": time.Now().Format(time.Kitchen)}
	return c.Render(http.StatusOK, "time", ctx)
}
