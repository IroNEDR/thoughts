package main

import (
	"embed"
	"errors"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"path/filepath"
)

var (
	templateFuncs = template.FuncMap{}
	//go:embed templates
	content embed.FS
)

type TemplateCache map[string]*template.Template

func CreatetemplateCache() (TemplateCache, error) {
	cache := TemplateCache{}
	// List all the pages that are in the "templates" folder of the content which end with "page.tmpl"
	pages, err := fs.Glob(content, "templates/*.page.tmpl")
	if err != nil {
		return nil, err
	}
	layouts, err := fs.Glob(content, "templates/layout.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		pageName := filepath.Base(page)
		ts, err := template.New(pageName).Funcs(templateFuncs).ParseFS(content, page)
		if err != nil {
			return nil, err
		}
		if len(layouts) == 1 {
			ts, err = ts.ParseFS(content, layouts[0])
			if err != nil {
				return nil, err
			}
		} else {
			return nil, errors.New("layout missing")
		}
		cache[pageName] = ts
	}
	return cache, nil
}

func Render(w http.ResponseWriter, tpl string) error {
	for k := range app.TemplCache {
		log.Println(k)
	}
	templ, ok := app.TemplCache[tpl]
	if !ok {
		return errors.New("template not found")
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	err := templ.Execute(w, "ekino")
	if err != nil {
		return err
	}

	return nil
}
