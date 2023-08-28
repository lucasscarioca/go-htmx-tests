package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name string
}

type Playlist struct {
	Title       string
	Description string
	CoverURL    string
	Categories  []string
}

type Tab struct {
	Title       string
	Author      string
	Description string
	Categories  []string
	CoverURL    string
	Tuning      string
	Content     string
	Public      bool
}

func Home(c echo.Context) error {
	ctx := map[string]any{
		"User": User{
			Name: "Lucas",
		},
		"PinnedPlaylists": []Playlist{
			{"Album1", "Album1 description", "", []string{"Guitar", "Pop"}},
			{"Album2", "Album2 description", "", []string{"Guitar", "Pop"}},
			{"Album3", "Album3 description", "", []string{"Guitar", "Pop"}},
			{"Album4", "Album4 description", "", []string{"Guitar", "Pop"}},
		},
		"StarredTabs": []Tab{
			{"Tab1", "Lucas", "Tab1 Description", []string{"Pop", "Guitar"}, "", "C", "test content", true},
			{"Tab2", "Lucas", "Tab2 Description", []string{"Pop", "Guitar"}, "", "C", "test content", true},
			{"Tab3", "Lucas", "Tab3 Description", []string{"Pop", "Guitar"}, "", "C", "test content", true},
			{"Tab4", "Lucas", "Tab4 Description", []string{"Pop", "Guitar"}, "", "C", "test content", true},
			{"Tab5", "Lucas", "Tab5 Description", []string{"Pop", "Guitar"}, "", "C", "test content", true},
			{"Tab6", "Lucas", "Tab6 Description", []string{"Pop", "Guitar"}, "", "C", "test content", true},
		},
	}
	return c.Render(http.StatusOK, "home", ctx)
}
