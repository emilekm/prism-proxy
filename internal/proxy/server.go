package proxy

import (
	"context"
	"strconv"

	"github.com/emilekm/go-prbf2/prism"
	"github.com/emilekm/prism-proxy/prismproxy"
)

func (p *Proxy) GetServerDetails(ctx context.Context, _ *prismproxy.Empty) (*prismproxy.ServerDetails, error) {
	details, err := p.c.Server.Details(ctx)
	if err != nil {
		return nil, err
	}

	return prismServerDetailsToProto(details), nil
}

func (p *Proxy) ServerDetailsUpdates(_ *prismproxy.Empty, conn prismproxy.Proxy_ServerDetailsUpdatesServer) error {
	ch, err := p.c.Server.DetailsUpdates(conn.Context())
	if err != nil {
		return err
	}
	defer p.c.Unsubscribe(ch)

	for {
		select {
		case <-conn.Context().Done():
			return nil
		case msg := <-ch:
			var details prism.ServerDetails
			err := prism.Unmarshal(msg.Body(), &details)
			if err != nil {
				return err
			}

			err = conn.Send(prismServerDetailsToProto(&details))
			if err != nil {
				return err
			}
		}
	}
}

func prismServerDetailsToProto(details *prism.ServerDetails) *prismproxy.ServerDetails {
	port, err := strconv.Atoi(details.Port)
	if err != nil {
		port = 0
	}

	return &prismproxy.ServerDetails{
		Name:           details.Name,
		Ip:             details.IP,
		Port:           int32(port),
		StartTime:      details.StartTime,
		RoundWarmup:    int32(details.RoundWarmup),
		RoundLength:    int32(details.RoundLength),
		MaxPlayers:     int32(details.MaxPlayers),
		Status:         details.Status,
		Map:            prismMapToProto(&details.Map),
		RoundStartTime: details.RoundStartTime,
		Players:        int32(details.Players),
		Team1:          details.Team1,
		Team2:          details.Team2,
		Tickets1:       int32(details.Tickets1),
		Tickets2:       int32(details.Tickets2),
		ConnectedUsers: details.ConnectedUsers,
	}
}

func prismMapToProto(details *prism.Map) *prismproxy.ServerMap {
	return &prismproxy.ServerMap{
		Name:  details.Name,
		Mode:  details.Mode,
		Layer: prismLayerToProto(details.Layer),
	}
}

func prismLayerToProto(layer prism.Layer) prismproxy.Layer {
	switch layer {
	case prism.LayerInfantry:
		return prismproxy.Layer_INF
	case prism.LayerAlternative:
		return prismproxy.Layer_ALT
	case prism.LayerStandard:
		return prismproxy.Layer_STD
	case prism.LayerLarge:
		return prismproxy.Layer_LRG
	case prism.LayerUnknown:
		fallthrough
	default:
		return prismproxy.Layer_None
	}
}
