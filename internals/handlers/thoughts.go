package handlers

import (
	"fmt"
	"net/http"

	"github.com/IroNEDR/thoughts/internals/config"
	"github.com/IroNEDR/thoughts/internals/renderer"
)

type ThoughtHandler interface {
	List(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type thoughtHandler struct {
	app      *config.AppConfig
	renderer renderer.Renderer
}

func NewThoughtHandler(app *config.AppConfig, renderer renderer.Renderer) ThoughtHandler {
	return &thoughtHandler{app, renderer}
}

func (th *thoughtHandler) List(w http.ResponseWriter, r *http.Request) {
	tmpl, err := th.renderer.LoadTemplate("home.page.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found"))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

func (th *thoughtHandler) Get(w http.ResponseWriter, r *http.Request) {
	tmpl, err := th.renderer.LoadTemplate("about.page.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Not found"))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

func (th *thoughtHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}

func (th *thoughtHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}

func (th *thoughtHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}
