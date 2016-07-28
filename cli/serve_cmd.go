package cli

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/citwild/wfe/api"
	"github.com/citwild/wfe/cli/internal/middleware"
	"github.com/citwild/wfe/service"
	"github.com/citwild/wfe/store/localstore"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"net"
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
	HTTPAddr  string `long:"http-addr" default:":8080" description:"HTTP listen address for app and gRPC API" env:"WFE_HTTP_ADDR"`
	HTTPSAddr string `long:"https-addr" default:":8443" description:"HTTPS (TLS) listen address for app and gRPC API" env:"WFE_HTTPS_ADDR"`

	CertFile string `long:"tls-cert" description:"certificate file for TLS" env:"WFE_TLS_CERT"`
	KeyFile  string `long:"tls-key" description:"key file for TLS" env:"WFE_TLS_KEY"`
}

func (c *ServeCmd) Execute(_ []string) error {
	// gRPC API
	grpcConfig := &grpcConfig{}
	grpcConfig.servers = service.NewServers()
	inj := middleware.NewInjector(grpcConfig.servers, localstore.NewStores())
	grpcConfig.opts = []grpc.ServerOption{grpc.UnaryInterceptor(inj.Inject)}

	// web app
	httpHandler := http.NewServeMux()
	httpHandler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Not yet implemented")
	})

	err := serveHTTP(c.HTTPAddr, grpcConfig, httpHandler)
	if err != nil {
		return err
	}

	useTLS := c.CertFile != "" || c.KeyFile != ""
	if useTLS {
		err = serveHTTPS(c.HTTPSAddr, grpcConfig, httpHandler, c.CertFile, c.KeyFile)
		if err != nil {
			return err
		}
	}

	select {}
}

func serveHTTP(addr string, grpcConfig *grpcConfig, httpHandler http.Handler) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	// main multiplexer
	mux := cmux.New(lis)

	// gRPC API
	grpcSrv := api.NewServer(grpcConfig.servers, grpcConfig.opts...)

	// web app
	httpSrv := &http.Server{}
	httpSrv.Addr = addr
	httpSrv.Handler = httpHandler

	// multiplex connection between gRPC API and app
	grpcLis := mux.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpLis := mux.Match(cmux.Any())

	log.Print("HTTP running on ", addr)
	go func() { log.Fatal(grpcSrv.Serve(grpcLis)) }()
	go func() { log.Fatal(httpSrv.Serve(httpLis)) }()
	go func() { log.Fatal(mux.Serve()) }()

	return nil
}

func serveHTTPS(addr string, grpcConfig *grpcConfig, httpHandler http.Handler, certFile string, keyFile string) error {
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
	grpcSrv := api.NewServer(grpcConfig.servers, grpcConfig.opts...)

	// multiplex connection between gRPC API and app
	srv.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcSrv.ServeHTTP(w, r)
		} else {
			httpHandler.ServeHTTP(w, r)
		}
	})

	log.Print("HTTPS running on ", addr)
	go func() { log.Fatal(srv.Serve(lis)) }()

	return nil
}

type grpcConfig struct {
	servers api.Servers
	opts    []grpc.ServerOption
}
