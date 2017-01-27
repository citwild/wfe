package app

import (
	"github.com/citwild/wfe/app/internal"
	"github.com/citwild/wfe/log"
	"net/http"
)

func NewHandler(r *Router) http.Handler {

	m := http.NewServeMux()

	for route, handler := range internal.Handlers {
		r.Get(route).Handler(handler)
		log.Info(route)
	}

	m.Handle("/", r)

	return m
}
