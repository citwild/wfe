package cli

import (
	"log"
	"sourcegraph.com/sourcegraph/go-flags"
)

var cmdInits []func(*flags.Parser)

func Main() error {
	log.SetFlags(0)

	parser := flags.NewNamedParser("wfe", flags.Default)
	parser.LongDescription = "wfe runs and manages a WFE instance."

	for _, init := range cmdInits {
		init(parser)
	}

	_, err := parser.Parse()
	return err
}
