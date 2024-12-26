package proxy

import (
	"github.com/emilekm/go-prbf2/prism"
	"github.com/emilekm/prism-proxy/prismproxy"
)

var _ prismproxy.ProxyServer = &Proxy{}

type Proxy struct {
	c *prism.Client

	prismproxy.UnimplementedProxyServer
}

func New(c *prism.Client) *Proxy {
	return &Proxy{c: c}
}
