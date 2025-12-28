package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"slices"
	"time"

	"github.com/emilekm/go-prbf2/prism"
	"github.com/emilekm/prism-proxy/internal/proxy"
	"github.com/emilekm/prism-proxy/prismproxy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var port = flag.String("port", "50051", "port to listen on")
var insecure = flag.Bool("insecure", false, "use insecure connection (no TLS)")
var certPath = flag.String("cert", "", "path to TLS certificate")
var keyPath = flag.String("key", "", "path to TLS key")
var prismAddr = flag.String("prism-addr", "", "address of the PRISM server ip:port")
var prismUser = flag.String("prism-user", "", "username for PRISM server")

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	flag.Parse()

	prismPass := os.Getenv("PRISM_PASS")

	if slices.Contains([]string{*prismAddr, *prismUser, prismPass}, "") {
		return fmt.Errorf("prism-addr, prism-user, and PRISM_PASS environment variable must be set")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", *port))
	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{}
	if !(*insecure) {
		if *certPath == "" || *keyPath == "" {
			return fmt.Errorf("cert and key paths must be provided in secure mode")
		}

		creds, err := credentials.NewServerTLSFromFile(*certPath, *keyPath)
		if err != nil {
			return err
		}
		opts = append(opts, grpc.Creds(creds))
	}

	gRPCserver := grpc.NewServer(opts...)

	prismClient, err := prism.Dial(*prismAddr)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = prismClient.Login(ctx, *prismUser, prismPass)
	if err != nil {
		return err
	}

	log.Printf("Connected and logged into PRISM at %s", *prismAddr)

	p := proxy.New(prismClient)

	prismproxy.RegisterProxyServer(gRPCserver, p)

	log.Printf("Starting proxy on port %s", *port)
	return gRPCserver.Serve(lis)
}
