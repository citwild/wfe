package tmpl

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	tmpldata "github.com/citwild/wfe/app/templates"
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/httputil"
)

var templates = map[string]*template.Template{}

func Load() {
	err := parseTemplates([][]string{{"layout.html"}})
	if err != nil {
		log.Fatal(err)
	}
}

func parseTemplates(sets [][]string) error {
	for _, set := range sets {
		tmpl := template.New("")

		for _, name := range set {
			f, err := tmpldata.Templates.Open("/" + name)
			if err != nil {
				log.Fatalf("Open template %s: %s", name, err)
			}
			t, err := ioutil.ReadAll(f)
			f.Close()
			if err != nil {
				log.Fatalf("Read template %s: %s", name, err)
			}
			_, err = tmpl.Parse(string(t))
			if err != nil {
				log.Fatalf("Parse template %s: %s", set, err)
			}
		}

		tmpl = tmpl.Lookup("ROOT")
		if tmpl == nil {
			return fmt.Errorf("ROOT template not found in %v", set)
		}
		templates[set[0]] = tmpl
	}
	return nil
}

func Execute(w http.ResponseWriter, r *http.Request, name string, status int, data interface{}) error {

	var bw httputil.ResponseBuffer

	bw.WriteHeader(status)
	if ct := bw.Header().Get("content-type"); ct == "" {
		bw.Header().Set("Content-Type", "text/html; charset=utf-8")
	}

	t := templates[name]
	if t == nil {
		return fmt.Errorf("Template %s not found", name)
	}

	err := t.Execute(&bw, data)
	if err != nil {
		return err
	}

	return bw.WriteTo(w)
}
