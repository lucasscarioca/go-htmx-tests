package controllers

import (
	"net/http"

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
	return c.Render(http.StatusOK, "home.html", todos)
}
