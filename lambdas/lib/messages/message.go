package messages

import (
	"encoding/json"
	"time"
)

type MessageOutput struct {
	Sender      *Member    `json:"sender"`
	Receiver    *Member    `json:"receiver"`
	Content     string     `json:"content"`
	ReceivedAt  *time.Time `json:"receivedAt"`
	MessageType string     `json:"type"`
}

func (output *MessageOutput) Encode() ([]byte, error) {
	return json.Marshal(output)
}

func NewMessageOutput(sender, receiver *Member, content string, receivedAt *time.Time, messageType string) *MessageOutput {
	return &MessageOutput{
		Sender:      sender,
		Receiver:    receiver,
		Content:     content,
		ReceivedAt:  receivedAt,
		MessageType: messageType,
	}
}
