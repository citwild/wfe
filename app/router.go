package app

import "github.com/gorilla/mux"

const (
	RouteUI = "ui"
)

type Router struct {
	mux.Router
}

func NewRouter(r *mux.Router) *Router {
	r.StrictSlash(true)

	r.PathPrefix("/").Methods("GET").Name(RouteUI)

	return &Router{*r}
}
