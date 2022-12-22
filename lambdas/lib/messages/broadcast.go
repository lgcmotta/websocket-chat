package messages

import "encoding/json"

type BroadcastMessageInput struct {
	Content json.RawMessage `json:"content"`
}

func (input *BroadcastMessageInput) Decode(message []byte) (*BroadcastMessageInput, error) {
	err := json.Unmarshal(message, input)
	return input, err
}
