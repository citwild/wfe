package tmpl

import (
	"html/template"
	"fmt"
	tmpldata "github.com/citwild/wfe/app/templates"
	"log"
	"io/ioutil"
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
				log.Fatalf("open template %s: %s", name, err)
			}
			t, err := ioutil.ReadAll(f)
			f.Close()
			if err != nil {
				log.Fatalf("read template %s: %s", name, err)
			}
			_, err = tmpl.Parse(string(t));
			if err != nil {
				log.Fatalf("parse template %s: %s", set, err)
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
