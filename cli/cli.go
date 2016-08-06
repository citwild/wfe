package cli

import (
	"sourcegraph.com/sourcegraph/go-flags"
)

var cmdInits []func(*flags.Parser)

var globalOpt struct {
	LogLevel string `long:"log-level" description:"upper log level to restrict log output to (debug, info, warn, error, panic, fatal)" default:"info" env:"WFE_LOG_LEVEL"`
}

func Main() error {
	parser := flags.NewNamedParser("wfe", flags.Default)
	parser.LongDescription = "wfe runs and manages a WFE instance."
	parser.AddGroup("Global options", "", &globalOpt)

	for _, init := range cmdInits {
		init(parser)
	}

	_, err := parser.Parse()
	return err
}
