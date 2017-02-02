package app

import (
	"net/http"
	"github.com/citwild/wfe/app/internal"
)

func NewHandler(r *Router) http.Handler {

	m := http.NewServeMux()

	for route, handler := range internal.Handlers {
		r.Get(route).Handler(handler)
	}

	m.Handle("/", r)

	return m
}
