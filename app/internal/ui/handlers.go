package ui

import (
	"net/http"

	"github.com/citwild/wfe/app"
	"github.com/citwild/wfe/app/internal"
	"github.com/citwild/wfe/app/internal/tmpl"
)

func init() {
	router.Get(routeTopLevel).HandlerFunc(serve)
	router.PathPrefix("/").Methods("GET").HandlerFunc(serve)

	internal.Handlers[app.RouteUI] = router
}

func serve(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, r, "layout.html", http.StatusOK, struct{}{})
}
