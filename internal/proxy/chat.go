package proxy

import (
	"github.com/emilekm/go-prbf2/prism"
	"github.com/emilekm/prism-proxy/prismproxy"
)

func (p *Proxy) ChatMessages(_ *prismproxy.Empty, stream prismproxy.Proxy_ChatMessagesServer) error {
	ch := p.c.Subscribe(prism.SubjectChat)
	defer p.c.Unsubscribe(ch)

	for {
		select {
		case <-stream.Context().Done():
			return nil
		case msg := <-ch:
			var chatMsgs prism.ChatMessages
			err := prism.Unmarshal(msg.Body(), &chatMsgs)
			if err != nil {
				return err
			}

			for _, chatMsg := range chatMsgs {
				err = stream.Send(prismChatMessageToProxy(&chatMsg))
				if err != nil {
					return err
				}
			}
		}
	}
}

func prismChatMessageToProxy(msg *prism.ChatMessage) *prismproxy.ChatMessage {
	return &prismproxy.ChatMessage{
		Type:       prismChatMessageTypeToProxy(msg.Type),
		Timestamp:  msg.Timestamp,
		Channel:    msg.Channel,
		PlayerName: msg.PlayerName,
		Content:    msg.Content,
	}
}

func prismChatMessageTypeToProxy(typ prism.ChatMessageType) prismproxy.ChatMessageType {
	switch typ {
	case prism.ChatMessageTypeOpfor:
		return prismproxy.ChatMessageType_OPFOR
	case prism.ChatMessageTypeBlufor:
		return prismproxy.ChatMessageType_BLUFOR
	case prism.ChatMessageTypeSquad:
		return prismproxy.ChatMessageType_SQUAD
	case prism.ChatMessageTypeServerMessage:
		return prismproxy.ChatMessageType_SERVER_MESSAGE
	case prism.ChatMessageTypeServer:
		return prismproxy.ChatMessageType_SERVER
	case prism.ChatMessageTypeResponse:
		return prismproxy.ChatMessageType_RESPONSE
	case prism.ChatMessageTypeAdminAlert:
		return prismproxy.ChatMessageType_ADMIN_ALERT
	default:
		return prismproxy.ChatMessageType_UNKNOWN
	}
}
