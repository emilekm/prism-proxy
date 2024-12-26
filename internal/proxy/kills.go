package proxy

import (
	"github.com/emilekm/go-prbf2/prism"
	"github.com/emilekm/prism-proxy/prismproxy"
)

func (p *Proxy) KillMessages(_ *prismproxy.Empty, stream prismproxy.Proxy_KillMessagesServer) error {
	ch := p.c.Subscribe(prism.SubjectKill)
	defer p.c.Unsubscribe(ch)

	for {
		select {
		case <-stream.Context().Done():
			return nil
		case msg := <-ch:
			var killMsg prism.KillMessages
			err := prism.Unmarshal(msg.Body(), &killMsg)
			if err != nil {
				return err
			}

			for _, kill := range killMsg {
				err = stream.Send(prismKillMessageToProxy(&kill))
				if err != nil {
					return err
				}
			}
		}
	}
}

func prismKillMessageToProxy(kill *prism.KillMessage) *prismproxy.KillMessage {
	return &prismproxy.KillMessage{
		IsTeamKill:   kill.IsTeamKill,
		Timestamp:    int64(kill.Timestamp),
		AttackerName: kill.AttackerName,
		VictimName:   kill.VictimName,
		Weapon:       kill.Weapon,
	}
}
