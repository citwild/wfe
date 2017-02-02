// +build generate

package main

import (
	"github.com/citwild/wfe/app/templates"
	"github.com/shurcooL/vfsgen"
	"log"
)

func main() {
	err := vfsgen.Generate(templates.Templates, vfsgen.Options{
		PackageName:  "templates",
		BuildTags:    "dist",
		VariableName: "Templates",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
