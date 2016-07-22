package cli

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/citwild/wfe/api"
	"github.com/citwild/wfe/cli/internal/middleware"
	"github.com/citwild/wfe/servers"
	"github.com/citwild/wfe/stores/localstores"
	"google.golang.org/grpc"
	"sourcegraph.com/sourcegraph/go-flags"
)

func init() {
	cmdInits = append(cmdInits, func(parser *flags.Parser) {
		_, err := parser.AddCommand("serve",
			"start web server",
			"Starts an HTTP server running the app and API.",
			&ServeCmd{})
		if err != nil {
			log.Fatal(err)
		}
	})
}

type ServeCmd struct {
	HTTPSAddr string `long:"https-addr" default:":8443" description:"HTTPS (TLS) listen address for app and gRPC API" env:"WFE_HTTPS_ADDR"`

	CertFile string `long:"tls-cert" description:"certificate file for TLS" env:"WFE_TLS_CERT" required:"yes"`
	KeyFile  string `long:"tls-key" description:"key file for TLS" env:"WFE_TLS_KEY" required:"yes"`
}

func (c *ServeCmd) Execute(_ []string) error {
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

	// main server
	srv := &http.Server{}
	srv.TLSConfig = config

	// gRPC API
	srvs := servers.NewServers()
	strs := localstores.NewLocalStores()
	inj := middleware.NewInjector(srvs, strs)
	grpcSrv := api.NewServer(srvs, grpc.UnaryInterceptor(inj.Inject))

	// web app
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Not yet implemented")
	})

	// multiplex connection between gRPC API and app
	srv.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcSrv.ServeHTTP(w, r)
		} else {
			mux.ServeHTTP(w, r)
		}
	})

	go func() { log.Fatal(srv.Serve(lis)) }()

	return nil
}
