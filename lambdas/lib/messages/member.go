package messages

import (
	"encoding/json"
	"fmt"
)

type Member struct {
	ConnectionId string `json:"connectionId"`
	Nickname     string `json:"nickname"`
}

func NewMember(connectionId, nickname string) *Member {
	return &Member{
		ConnectionId: connectionId,
		Nickname:     nickname,
	}
}

func (member *Member) GetJoiningMessage() string {
	return fmt.Sprintf("%s just joined the chat", member.Nickname)
}

func (member *Member) GetLeavingMessage() string {
	return fmt.Sprintf("%s has left the chat", member.Nickname)
}

type ConnectedMembers struct {
	Members []*Member `json:"members"`
}

func (connected *ConnectedMembers) Encode() ([]byte, error) {
	return json.Marshal(connected)
}
