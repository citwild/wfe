package cli

import (
	"log"
	"os"
	"sourcegraph.com/sourcegraph/go-flags"
)

var cmdInits []func(*flags.Parser)

func Main() error {
	log.SetFlags(0)

	name := os.Args[0]
	parser := flags.NewNamedParser(name, flags.Default)
	parser.LongDescription = name + " runs and manages a WFE instance."

	for _, init := range cmdInits {
		init(parser)
	}

	_, err := parser.Parse()
	return err
}
