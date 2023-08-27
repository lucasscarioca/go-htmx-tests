package views

import (
	"errors"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

var templatesMap = map[string][]string{
	// Pages
	"home": {
		"internal/views/base.html",
		"internal/views/components/navbar.html",
		"internal/views/pages/home.html",
	},
	"pageNotFound": {
		"internal/views/base.html",
		"internal/views/components/navbar.html",
		"internal/views/pages/404.html",
	},
	// HTMX Fragments
	"time": {
		"internal/views/components/time.html",
	},
}

type TemplateRegistry struct {
	templates map[string]*template.Template
}

func WithTemplateRegistry() *TemplateRegistry {
	return &TemplateRegistry{
		templates: registerTemplates(),
	}
}

// Implement e.Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data any, c echo.Context) error {
	template, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found: " + name)
		return err
	}
	base := template.Lookup("base.html")
	if base == nil {
		return template.ExecuteTemplate(w, name, data)
	}
	return template.ExecuteTemplate(w, "base.html", data)
}

func registerTemplates() map[string]*template.Template {
	templates := make(map[string]*template.Template)

	for key, value := range templatesMap {
		templates[key] = newTemplate(value...)
	}

	return templates
}

func newTemplate(files ...string) *template.Template {
	return template.Must(template.New("").ParseFiles(files...))
}
