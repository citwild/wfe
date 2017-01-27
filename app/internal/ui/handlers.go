package ui

import (
	"fmt"
	"github.com/citwild/wfe/app"
	"github.com/citwild/wfe/app/internal"
	"net/http"
)

func init() {
	router.Get(routeTopLevel).HandlerFunc(serve)
	router.PathPrefix("/").Methods("GET").HandlerFunc(serve)

	internal.Handlers[app.RouteUI] = router
}

func serve(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!")
}
