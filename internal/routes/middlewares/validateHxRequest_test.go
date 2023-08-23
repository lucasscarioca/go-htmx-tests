package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestInvalidHxRequest(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/hx/time", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := ValidateHxRequest(func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	err := h(c)
	code := http.StatusInternalServerError
	if err != nil {
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}
	}
	assert.Equal(t, http.StatusBadRequest, code)
}
