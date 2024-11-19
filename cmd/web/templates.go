package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"time"

	"github.com/MohitPanchariya/Snippet-Box/internal/models"
	"github.com/MohitPanchariya/Snippet-Box/ui"
)

// structure to hold dynamic data passed to html templates
type templateData struct {
	CurrentYear     int
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
	Flash           string
	IsAuthenticated bool
	Form            any
	CSRFToken       string
}

// Cache templates at application start up
func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		// slice containing filepath patterns for the templates we want
		// to parse.
		patters := []string{
			"html/base.tmpl.html",
			"html/partials/*.tmpl.html",
			page,
		}
		// Parse the partials and associate them with the base template
		ts, err := template.New(name).ParseFS(ui.Files, patters...)
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}

	return cache, nil
}

func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.UTC().Format("02 Jan 2006 at 15:04")
}
