package cli

import (
	"crypto/tls"
	"fmt"
	"github.com/citwild/wfe/api"
	"github.com/citwild/wfe/services"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"sourcegraph.com/sourcegraph/go-flags"
	"strings"
)

func init() {
	cmdInits = append(cmdInits, func(parser *flags.Parser) {
		_, err := parser.AddCommand("serve",
			"start web server",
			"Starts an HTTP server running the app and API.",
			&serveCmd{})
		if err != nil {
			log.Fatal(err)
		}
	})
}

type ServeCmd struct {
	HTTPSAddr string `long:"https-addr" default:":3443" description:"HTTPS (TLS) listen address for app and gRPC API" env:"WFE_HTTPS_ADDR"`

	CertFile string `long:"tls-cert" description:"certificate file for TLS" env:"WFE_TLS_CERT" required:"yes"`
	KeyFile  string `long:"tls-key" description:"key file for TLS" env:"WFE_TLS_KEY" required:"yes"`
}

type serveCmd ServeCmd

func (c *serveCmd) Execute(_ []string) error {
	err := serveHTTPS(c.HTTPSAddr, c.CertFile, c.KeyFile)
	if err != nil {
		return err
	}

	select {}
}

func serveHTTPS(addr string, certFile string, keyFile string) error {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return err
	}

	config := &tls.Config{Certificates: []tls.Certificate{cert}}

	lis, err := tls.Listen("tcp", addr, config)
	if err != nil {
		return err
	}

	srv := &http.Server{}
	srv.TLSConfig = config

	grpcSrv := grpc.NewServer()
	api.RegisterAccountsServer(grpcSrv, services.Accounts)

	// multiplex connection between app and gRPC API
	srv.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcSrv.ServeHTTP(w, r)
		} else {
			fmt.Fprintln(w, "Not yet implemented")
		}
	})

	go func() { log.Fatal(srv.Serve(lis)) }()

	return nil
}
