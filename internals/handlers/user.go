package handlers

import (
	"fmt"
	"net/http"
)

type UserHandler interface {
	Me(w http.ResponseWriter, r *http.Request)
}

type userHandler struct{}

func NewUserHandler() UserHandler {
	return &userHandler{}
}

func (uh *userHandler) Me(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	fmt.Println(w, "ekin")
}
