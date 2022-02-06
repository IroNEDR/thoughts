package renderer

import (
	"errors"
	"html/template"
	"log"
	"path/filepath"

	"github.com/IroNEDR/thoughts/internals/config"
)

var (
	templateFuncs = template.FuncMap{}
)

type Renderer interface {
	CreateTemplateCache() (config.TemplateCache, error)
	LoadTemplate(tmpl string) (*template.Template, error)
}

type renderer struct {
	app *config.AppConfig
}

func NewRenderer(app *config.AppConfig) Renderer {
	return &renderer{app}
}

func (r *renderer) CreateTemplateCache() (config.TemplateCache, error) {
	cache := config.TemplateCache{}
	// List all the pages that are in the "templates" folder of the content which end with "page.tmpl"
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return nil, err
	}
	layouts, err := filepath.Glob("templates/*.layout.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		pageName := filepath.Base(page)
		ts, err := template.New(pageName).Funcs(templateFuncs).ParseFiles(page)
		if err != nil {
			return nil, err
		}
		if len(layouts) > 0 {
			for _, layout := range layouts {
				ts, err = ts.ParseFiles(layout)
				if err != nil {
					return nil, err
				}
			}
		} else {
			return nil, errors.New("layouts missing")
		}
		cache[pageName] = ts
	}
	return cache, nil
}

func (r *renderer) LoadTemplate(tpl string) (*template.Template, error) {
	for k := range r.app.TemplCache {
		log.Println(k)
	}
	templ, ok := r.app.TemplCache[tpl]
	if !ok {
		return nil, errors.New("template not found")
	}
	return templ, nil
}
