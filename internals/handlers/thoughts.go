package handlers

import (
	"fmt"
	"net/http"

	"github.com/IroNEDR/thoughts/internals/config"
	"github.com/IroNEDR/thoughts/internals/renderer"
)

type ThoughtHandler struct {
	app      *config.AppConfig
	renderer *renderer.Renderer
}

func NewThoughtHandler(app *config.AppConfig, renderer *renderer.Renderer) *ThoughtHandler {
	return &ThoughtHandler{app, renderer}
}

func (th *ThoughtHandler) List(w http.ResponseWriter, r *http.Request) {
	tmpl, err := th.renderer.LoadTemplate("home.page.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found"))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

func (th *ThoughtHandler) Get(w http.ResponseWriter, r *http.Request) {
	tmpl, err := th.renderer.LoadTemplate("about.page.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Not found"))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

func (th *ThoughtHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}

func (th *ThoughtHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}

func (th *ThoughtHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}
