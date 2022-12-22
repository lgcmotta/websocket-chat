package messages

import "fmt"

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

func (member *Member) GetJoiningMessage() []byte {
	message := fmt.Sprintf(`"%s just joined the chat"`, member.Nickname)
	return []byte(message)
}

func (member *Member) GetLeavingMessage() []byte {
	message := fmt.Sprintf(`"%s has left the chat"`, member.Nickname)
	return []byte(message)
}
