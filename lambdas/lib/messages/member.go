package messages

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
