package messages

import (
	"encoding/json"
)

type DirectMessageInput struct {
	Receivers []string `json:"receivers"`
	Content   string   `json:"content"`
}

func (input *DirectMessageInput) Decode(message []byte) (*DirectMessageInput, error) {
	err := json.Unmarshal(message, input)
	return input, err
}
