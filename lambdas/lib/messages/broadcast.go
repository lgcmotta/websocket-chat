package messages

import (
	"encoding/json"
	"time"
)

type BroadcastMessageInput struct {
	Content string `json:"content"`
}

func (input *BroadcastMessageInput) Decode(message []byte) (*BroadcastMessageInput, error) {
	err := json.Unmarshal(message, input)
	return input, err
}

type BroadcastMessageOutput struct {
	Sender     *Member    `json:"sender"`
	Receivers  []*Member  `json:"receivers"`
	Content    string     `json:"content"`
	ReceivedAt *time.Time `json:"receivedAt"`
}

func (output *BroadcastMessageOutput) Encode() ([]byte, error) {
	return json.Marshal(output)
}

func NewBroadcastMessageOutput(sender *Member, receivers []*Member, content string, receivedAt *time.Time) *BroadcastMessageOutput {
	return &BroadcastMessageOutput{
		Sender:     sender,
		Receivers:  receivers,
		Content:    content,
		ReceivedAt: receivedAt,
	}
}
