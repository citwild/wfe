package main

import (
	"flag"
	"fmt"
	"os"
	"net/http"
	"log"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, `wfe runs and manages a WFE instance.

Usage:

	wfe [options] command [arguments]

The commands are:
`)
		for _, c := range commands {
			fmt.Fprintf(os.Stderr, "	%-12s %s\n", c.name, c.description)
		}
		fmt.Fprintln(os.Stderr, `
Use "wfe command -h" for more information about a command.

The options are:
`)
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
	}

	log.SetFlags(0)

	cmd := flag.Arg(0)
	for _, c := range commands {
		if c.name == cmd {
			c.run(flag.Args()[1:])
			return
		}
	}

	fmt.Fprintf(os.Stderr, "unknown command %q\n", cmd)
	fmt.Fprintln(os.Stderr, `Run "wfe -h" for usage.`)
	os.Exit(1)
}

type command struct {
	name        string
	description string
	run         func(args []string)
}

var commands = []command{
	{"serve", "start web server", serve},
}

func serve(args []string) {
	fs := flag.NewFlagSet("serve", flag.ExitOnError)
	httpAddr := fs.String("http", ":5000", "HTTP service address")
	fs.Usage = func() {
		fmt.Fprintln(os.Stderr, `usage: wfe serve [options]

Starts the web server that servers the API.

The options are:
`)
		fs.PrintDefaults()
		os.Exit(1)
	}

	fs.Parse(args)
	if fs.NArg() != 0 {
		fs.Usage()
	}

	m := http.NewServeMux()

	log.Print("Listening on ", *httpAddr)
	err := http.ListenAndServe(*httpAddr, m)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}