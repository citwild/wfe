package cli

import (
	"log"
	"sourcegraph.com/sourcegraph/go-flags"
)

var cmdInits []func(*flags.Parser)

func Main(args []string) error {
	log.SetFlags(0)

	name := args[0]
	parser := flags.NewNamedParser(name, flags.Default)
	parser.LongDescription = name + " runs and manages a WFE instance."

	for _, init := range cmdInits {
		init(parser)
	}

	_, err := parser.Parse()
	return err
}
