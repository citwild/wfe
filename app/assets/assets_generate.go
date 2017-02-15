// +build generate

package main

import (
	"github.com/citwild/wfe/app/assets"
	"github.com/shurcooL/vfsgen"
	"log"
)

func main() {
	err := vfsgen.Generate(assets.Assets, vfsgen.Options{
		PackageName:  "assets",
		VariableName: "Assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
