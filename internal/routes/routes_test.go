package routes

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/lucasscarioca/music-stash/internal/routes/middlewares"
	"github.com/stretchr/testify/assert"
)

func init() {
	if err := os.Chdir("../.."); err != nil {
		panic(err)
	}
}

func setupServerTest() *echo.Echo {
	e := echo.New()
	middlewares.Mount(e)
	Mount(e)
	return e
}

func TestIfServerIsRunning(t *testing.T) {
	// Setup
	e := setupServerTest()
	req := httptest.NewRequest(http.MethodGet, "/ping", strings.NewReader(""))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Assertions
	if assert.NoError(t, pingHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "pong", rec.Body.String())
	}
}
