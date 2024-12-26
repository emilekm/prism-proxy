package proxy

import (
	"context"

	"github.com/emilekm/go-prbf2/prism"
	"github.com/emilekm/prism-proxy/prismproxy"
)

func (p *Proxy) GetGameplayDetails(ctx context.Context, _ *prismproxy.Empty) (*prismproxy.GameplayDetails, error) {
	details, err := p.c.Gameplay.Details(ctx)
	if err != nil {
		return nil, err
	}

	return prismGameplayDetailsToProto(details), nil
}

func (p *Proxy) GameplayDetailsUpdates(_ *prismproxy.Empty, stream prismproxy.Proxy_GameplayDetailsUpdatesServer) error {
	ch, err := p.c.Gameplay.DetailsUpdates(stream.Context())
	if err != nil {
		return err
	}
	defer p.c.Unsubscribe(ch)

	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		case msg := <-ch:
			var details *prism.GameplayDetails
			err = prism.Unmarshal(msg.Body(), &details)
			if err != nil {
				return err
			}

			err := stream.Send(prismGameplayDetailsToProto(details))
			if err != nil {
				return err
			}
		}
	}
}

func prismGameplayDetailsToProto(details *prism.GameplayDetails) *prismproxy.GameplayDetails {
	return &prismproxy.GameplayDetails{
		ControlPoints: prismControlPointsToProto(details.ControlPoints),
		Fobs:          prismFobsToProto(details.Fobs),
		Rallies:       prismRalliesToProto(details.Rallies),
	}
}

func prismControlPointsToProto(controlPoints []prism.ControlPoint) []*prismproxy.ControlPoint {
	var cps []*prismproxy.ControlPoint

	for _, cp := range controlPoints {
		cps = append(cps, &prismproxy.ControlPoint{
			Id:   cp.ID,
			Team: int32(cp.Team),
		})
	}

	return cps
}

func prismFobsToProto(fobs []prism.Fob) []*prismproxy.Fob {
	var f []*prismproxy.Fob

	for _, fob := range fobs {
		f = append(f, &prismproxy.Fob{
			Position: fob.Position,
			Team:     int32(fob.Team),
		})
	}

	return f
}

func prismRalliesToProto(rallies []prism.Rally) []*prismproxy.Rally {
	var r []*prismproxy.Rally

	for _, rally := range rallies {
		r = append(r, &prismproxy.Rally{
			Position: rally.Position,
			Team:     int32(rally.Team),
			Squad:    int32(rally.Squad),
		})
	}

	return r
}
