package main

import (
	"context"
	"flag"
	"log"
	"net"
	"time"

	"github.com/emilekm/go-prbf2/prism"
	"github.com/emilekm/prism-proxy/internal/config"
	"github.com/emilekm/prism-proxy/internal/proxy"
	"github.com/emilekm/prism-proxy/prismproxy"
	"google.golang.org/grpc"
)

var configPath = flag.String("config", "config.yaml", "path to config file")

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	flag.Parse()

	cfg, err := config.New(*configPath)
	if err != nil {
		return err
	}

	lis, err := net.Listen("tcp", ":7777")
	if err != nil {
		return err
	}

	gRPCserver := grpc.NewServer()

	prismClient, err := prism.Dial(net.JoinHostPort(cfg.PRISM.Host, cfg.PRISM.Port))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = prismClient.Login(ctx, cfg.PRISM.User, cfg.PRISM.Pass)
	if err != nil {
		return err
	}

	log.Printf("Connected and logged into PRISM at %s:%s", cfg.PRISM.Host, cfg.PRISM.Port)

	p := proxy.New(prismClient)

	prismproxy.RegisterProxyServer(gRPCserver, p)

	return gRPCserver.Serve(lis)
}
