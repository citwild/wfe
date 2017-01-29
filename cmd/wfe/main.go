package main

import (
	"github.com/citwild/wfe/cli"
	"os"

	_ "github.com/citwild/wfe/app/cmd"
)

func main() {
	err := cli.Main()
	if err != nil {
		os.Exit(1)
	}
}
