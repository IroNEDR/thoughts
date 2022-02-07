package handlers

import (
	"net/http"

	"github.com/IroNEDR/thoughts/internals/config"
	"github.com/IroNEDR/thoughts/internals/renderer"
)

type LandingHandler struct {
	app      *config.AppConfig
	renderer *renderer.Renderer
}

func NewLandingHandler(app *config.AppConfig, rd *renderer.Renderer) *LandingHandler {
	return &LandingHandler{app, rd}
}

func (lh *LandingHandler) Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := lh.renderer.LoadTemplate("index.page.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found"))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}
