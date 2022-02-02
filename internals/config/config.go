package config

import (
	"html/template"
)

type TemplateCache map[string]*template.Template

type AppConfig struct {
	TemplCache TemplateCache
	CSRFkey    []byte
}

type Environment struct {
	CSRF_KEY string
}
