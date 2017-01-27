package ui

import (
	"github.com/gorilla/mux"
	"strings"
)

const (
	routeTopLevel = "toplevel"
)

var router = newRouter()

func newRouter() *mux.Router {
	r := mux.NewRouter()

	r.StrictSlash(true)

	topLevel := []string{
		"about",
		"login",
	}
	r.Path("/{Path:(?:" + strings.Join(topLevel, "|") + ")}").Methods("GET").Name(routeTopLevel)

	return r
}
