package handlers

import (
	"fmt"
	"net/http"
)

type Category struct {
	Name        string
	Description string
	Icon        string
}

type CategoryHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

type categoryHandler struct {
}

func NewCategoryHandler() CategoryHandler {
	return &categoryHandler{}
}

func (ch *categoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	fmt.Println(w, "Category")
}

func (ch *categoryHandler) List(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Println(w, "[a,b,c]")
}
