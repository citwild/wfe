package templates

import (
	"go/build"
	"log"

	"github.com/shurcooL/httpfs/filter"
	"net/http"
)

func getDir(importPath string) string {
	p, err := build.Import(importPath, "", build.FindOnly)
	if err != nil {
		log.Fatalln(err)
	}
	return p.Dir
}

var Templates = filter.Skip(
	http.Dir(getDir("github.com/citwild/wfe/app/templates")),
	filter.FilesWithExtensions(".go"),
)
