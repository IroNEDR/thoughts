package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

type TemplateCache map[string]*template.Template

type AppConfig struct {
	TemplCache     TemplateCache
	SessionManager *scs.SessionManager
	Env            string
	CSRFkey        []byte
	IsProd         bool
}

type Environment struct {
	CSRF_KEY string
}
