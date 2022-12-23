package messages

import (
	"encoding/json"
	"time"
)

type MessageOutput struct {
	Sender     *Member         `json:"sender"`
	Receiver   *Member         `json:"receiver"`
	Content    json.RawMessage `json:"content"`
	ReceivedAt *time.Time      `json:"receivedAt"`
	Type       MessageType     `json:"type"`
}

func (output *MessageOutput) Encode() ([]byte, error) {
	return json.Marshal(output)
}

func NewMessageOutput(sender, receiver *Member, content json.RawMessage, receivedAt *time.Time, messageType MessageType) *MessageOutput {
	return &MessageOutput{
		Sender:     sender,
		Receiver:   receiver,
		Content:    content,
		ReceivedAt: receivedAt,
	}
}
