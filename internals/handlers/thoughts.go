package handlers

import (
	"fmt"
	"net/http"

	"github.com/IroNEDR/thoughts/internals/config"
	"github.com/IroNEDR/thoughts/internals/renderer"
)

type Thought struct {
	CreatedBy   User       `json:"created_by"`
	Name        string     `json:"name"`
	Description string     `json:"description,omitempty"`
	Categories  []Category `json:"categories"`
	Tags        []string   `json:"tags,omitempty"`
	Public      bool       `json:"public"`
}

type ThoughtHandler interface {
	List(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type thoughtHandler struct {
	app      *config.AppConfig
	renderer *renderer.Renderer
}

func NewThoughtHandler(app *config.AppConfig, renderer *renderer.Renderer) ThoughtHandler {
	return thoughtHandler{app, renderer}
}

func (th thoughtHandler) List(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Println("test")
	fmt.Fprint(w, "[a,b,c]")
}

func (th thoughtHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "a")
}

func (th thoughtHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}

func (th thoughtHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}

func (th thoughtHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}
