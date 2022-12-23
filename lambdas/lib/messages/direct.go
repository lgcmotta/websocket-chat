package messages

import (
	"encoding/json"
)

type DirectMessageInput struct {
	Receiver string `json:"receiver"`
	Content  string `json:"content"`
}

func (input *DirectMessageInput) Decode(message []byte) (*DirectMessageInput, error) {
	err := json.Unmarshal(message, input)
	return input, err
}
