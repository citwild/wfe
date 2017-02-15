package assets

import (
	"github.com/gorilla/mux"
	"net/http"
	"path"
)

func NewHandler(r *mux.Router) http.Handler {
	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := path.Clean(r.URL.Path)

		f, err := Assets.Open(name)
		if err != nil {
			http.Error(w, "Failed to open asset", http.StatusInternalServerError)
			return
		}
		defer f.Close()

		fi, _ := f.Stat()

		http.ServeContent(w, r, fi.Name(), fi.ModTime(), f)
	})

	r.PathPrefix("/").Handler(handler)

	return r
}
