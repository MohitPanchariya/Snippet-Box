package main

import (
	"html/template"
	"path/filepath"

	"github.com/MohitPanchariya/Snippet-Box/internal/models"
)

// structure to hold dynamic data passed to html templates
type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

// Cache templates at application start up
func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		// Parse the base template
		ts, err := template.ParseFiles("./ui/html/base.tmpl.html")
		if err != nil {
			return nil, err
		}
		// Parse the partials and associate them with the base template
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl.html")
		if err != nil {
			return nil, err
		}

		// Parse the page and associate it with the base template
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
