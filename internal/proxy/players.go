package proxy

import (
	"context"

	"github.com/emilekm/go-prbf2/prism"
	"github.com/emilekm/prism-proxy/prismproxy"
)

func (p *Proxy) GetPlayers(ctx context.Context, _ *prismproxy.Empty) (*prismproxy.Players, error) {
	players, err := p.c.Players.List(ctx)
	if err != nil {
		return nil, err
	}

	return prismPlayersToProto(players), nil
}

func (p *Proxy) PlayersUpdates(_ *prismproxy.Empty, stream prismproxy.Proxy_PlayersUpdatesServer) error {
	ch, err := p.c.Players.ListUpdates(stream.Context())
	if err != nil {
		return err
	}
	defer p.c.Unsubscribe(ch)

	for {
		select {
		case <-stream.Context().Done():
			return nil
		case msg := <-ch:
			var updates prism.UpdatePlayers
			err = prism.Unmarshal(msg.Body(), &updates)
			if err != nil {
				return err
			}

			err = stream.Send(prismPlayersUpdatesToProto(updates))
			if err != nil {
				return err
			}
		}
	}
}

func (p *Proxy) PlayerLeaveUpdates(_ *prismproxy.Empty, stream prismproxy.Proxy_PlayerLeaveUpdatesServer) error {
	ch, err := p.c.Players.PlayerLeaveUpdates(stream.Context())
	if err != nil {
		return err
	}
	defer p.c.Unsubscribe(ch)

	for {
		select {
		case <-stream.Context().Done():
			return nil
		case msg := <-ch:
			var name string

			err = prism.Unmarshal(msg.Body(), &name)
			if err != nil {
				return err
			}

			err = stream.Send(&prismproxy.PlayerLeave{Name: name})
		}
	}
}

func prismPlayersToProto(players []prism.FullPlayer) *prismproxy.Players {
	var ps []*prismproxy.Player

	for _, p := range players {
		ps = append(ps, &prismproxy.Player{
			Name:          p.Name,
			IsAiPlayer:    p.IsAIPlayer,
			Hash:          p.Hash,
			Ip:            p.IP,
			ProfileId:     p.ProfileID,
			Index:         int32(p.Index),
			JoinTimestamp: p.JoinTimestamp,
			PlayerDetails: prismPlayerDetailsToProto(p.PlayerDetails),
		})
	}

	return &prismproxy.Players{Players: ps}
}

func prismPlayersUpdatesToProto(updates prism.UpdatePlayers) *prismproxy.PlayersUpdate {
	var ps []*prismproxy.PlayerUpdate

	for _, p := range updates {
		if p.Full != nil {
			ps = append(ps, &prismproxy.PlayerUpdate{
				Partial: false,
				Player: &prismproxy.Player{
					Name:          p.Full.Name,
					IsAiPlayer:    p.Full.IsAIPlayer,
					Hash:          p.Full.Hash,
					Ip:            p.Full.IP,
					ProfileId:     p.Full.ProfileID,
					Index:         int32(p.Full.Index),
					JoinTimestamp: p.Full.JoinTimestamp,
					PlayerDetails: prismPlayerDetailsToProto(p.Full.PlayerDetails),
				},
			})
		} else {
			ps = append(ps, &prismproxy.PlayerUpdate{
				Partial: true,
				Player: &prismproxy.Player{
					Name:          p.Update.Name,
					Index:         int32(p.Update.Index),
					PlayerDetails: prismPlayerDetailsToProto(p.Update.PlayerDetails),
				},
			})
		}
	}

	return &prismproxy.PlayersUpdate{Players: ps}
}

func prismPlayerDetailsToProto(details prism.PlayerDetails) *prismproxy.PlayerDetails {
	return &prismproxy.PlayerDetails{
		Team:          int32(details.Team),
		Squad:         details.Squad,
		Kit:           details.Kit,
		Vehicle:       details.Vehicle,
		Score:         int32(details.Score),
		ScoreTeamwork: int32(details.ScoreTeamwork),
		Kills:         int32(details.Kills),
		Teamkills:     int32(details.Teamkills),
		Deaths:        int32(details.Deaths),
		Valid:         details.Valid,
		Ping:          int32(details.Ping),
		Idle:          details.Idle,
		Alive:         details.Alive,
		Joining:       details.Joining,
		Position:      details.Position,
		Rotation:      details.Rotation,
	}
}
