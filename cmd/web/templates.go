package main

import (
	"html/template"
	"path/filepath"

	"github.com/isuquo/snippetbox/internal/models"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	// Initialize a new map to act as the cache.
	cache := map[string]*template.Template{}

	// Use the filepath.Glob() function to get a slice of all filepaths with the
	// extension '.html'. This matches the naming pattern for our HTML page
	// templates.
	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	// Loop through the pages one-by-one.
	for _, page := range pages {
		// Extract the file name (like 'index.html') from the full file path
		name := filepath.Base(page)

		files := []string{
			"./ui/html/layouts/main.html",
			"./ui/html/partials/nav.html",
			page,
		}

		// Parse the files into a template set
		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}