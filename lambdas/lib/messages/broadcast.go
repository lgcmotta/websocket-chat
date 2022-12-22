package messages

import "encoding/json"

type BroadcastMessageInput struct {
	Content json.RawMessage `json:"content"`
}

func (input *BroadcastMessageInput) Decode(message []byte) (*BroadcastMessageInput, error) {
	err := json.Unmarshal(message, input)
	return input, err
}

type BroadcastMessageOutput struct {
	Sender    *Member         `json:"sender"`
	Receivers []*Member       `json:"receivers"`
	Content   json.RawMessage `json:"content"`
}

func (output *BroadcastMessageOutput) Encode() ([]byte, error) {
	return json.Marshal(output)
}

func NewBroadcastMessageOutput(sender *Member, receivers []*Member, content json.RawMessage) *BroadcastMessageOutput {
	return &BroadcastMessageOutput{
		Sender:    sender,
		Receivers: receivers,
		Content:   content,
	}
}
