// +build !dist

package assets

import (
	"github.com/shurcooL/httpfs/filter"
	"go/build"
	"log"
	"net/http"
)

func getDir(importPath string) string {
	p, err := build.Import(importPath, "", build.FindOnly)
	if err != nil {
		log.Fatalln(err)
	}
	return p.Dir
}

var Assets = filter.Skip(
	http.Dir(getDir("github.com/citwild/wfe/app/assets")),
	filter.FilesWithExtensions(".go", ".html"),
)
