package main

import (
	"html/template"
	"path/filepath"

	"github.com/vladfreishmidt/notenow/internal/models"
)

type templateData struct {
	Note  *models.Note
	Notes []*models.Note
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.gohtml")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		files := []string{
			"./ui/html/partials/nav.gohtml",
			"./ui/html/base.gohtml",
			page,
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
