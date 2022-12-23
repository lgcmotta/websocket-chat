package messages

type MessageType string

const (
	Broadcast MessageType = "broadcast"
	Direct    MessageType = "direct"
)
