package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, `wfe runs and manages a WFE instance.

Usage:

	wfe [options] command [arguments]

The commands are:
`)
		for _, c := range subcommands {
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
	for _, c := range subcommands {
		if c.name == cmd {
			c.run(flag.Args()[1:])
			return
		}
	}

	fmt.Fprintf(os.Stderr, "unknown subcommand %q\n", cmd)
	fmt.Fprintln(os.Stderr, `Run "wfe -h" for usage.`)
	os.Exit(1)
}

type subcommand struct {
	name        string
	description string
	run         func(args []string)
}

var subcommands = []subcommand{
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

	log.Print("Listening on ", *httpAddr)
	lis, err := net.Listen("tcp", *httpAddr)
	if err != nil {
		log.Fatal("Listen:", err)
	}

	s := grpc.NewServer()
	s.Serve(lis)
}
