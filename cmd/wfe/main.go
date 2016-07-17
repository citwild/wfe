package main

import (
	"github.com/citwild/wfe/cli"
	"os"
)

func main() {
	err := cli.Main(os.Args)
	if err != nil {
		os.Exit(1)
	}
}
