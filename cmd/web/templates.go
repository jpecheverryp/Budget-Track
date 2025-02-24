package main

import (
	"html/template"
	"io/fs"
	"path/filepath"

	"github.com/jpecheverryp/budget-track/internal/repository"
	"github.com/jpecheverryp/budget-track/ui"
)

type templateData struct {
	Accounts []repository.Account
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			"html/base.html",
			"html/partials/nav.html",
			page,
		}

		ts, err := template.New(name).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts

	}

	return cache, nil
}
