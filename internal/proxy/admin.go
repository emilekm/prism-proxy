package proxy

import (
	"context"

	"github.com/emilekm/prism-proxy/prismproxy"
)

func (p *Proxy) APIAdmin(ctx context.Context, req *prismproxy.APIAdminReq) (*prismproxy.APIAdminResp, error) {
	output, err := p.c.Admin.APIAdmin(ctx, req.GetCommand())
	if err != nil {
		return nil, err
	}

	return &prismproxy.APIAdminResp{Response: output}, nil
}

func (p *Proxy) RACommand(ctx context.Context, req *prismproxy.RACommandReq) (*prismproxy.RACommandResp, error) {
	output, err := p.c.Admin.RACommand(ctx, req.GetCommand())
	if err != nil {
		return nil, err
	}

	return &prismproxy.RACommandResp{
		Topic:   output.Topic,
		Content: output.Content,
	}, nil
}
