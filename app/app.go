package app

import (
	"github.com/citwild/wfe/app/internal"
	"net/http"
)

func NewHandler(r *Router) http.Handler {

	m := http.NewServeMux()

	for route, handler := range internal.Handlers {
		r.Get(route).Handler(handler)
	}

	m.Handle("/", r)

	return m
}
