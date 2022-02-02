package renderer

import (
	"embed"
	"errors"
	"html/template"
	"io/fs"
	"log"
	"path/filepath"

	"github.com/IroNEDR/thoughts/internals/config"
)

var (
	templateFuncs = template.FuncMap{}
	//go:embed templates
	content embed.FS
)

type Renderer interface {
	CreateTemplateCache() (config.TemplateCache, error)
	Render(tmpl string) (*template.Template, error)
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

func (r *renderer) Render(tpl string) (*template.Template, error) {
	for k := range r.app.TemplCache {
		log.Println(k)
	}
	templ, ok := r.app.TemplCache[tpl]
	if !ok {
		return nil, errors.New("template not found")
	}

	// w.Header().Set("Content-Type", "text/html")
	// w.WriteHeader(http.StatusOK)
	// err := templ.Execute(w, "ekino")
	// if err != nil {
	// 	return err
	// }

	return templ, nil
}
