package messages

import "encoding/json"

type DirectMessageInput struct {
	Receiver string          `json:"receiver"`
	Content  json.RawMessage `json:"content"`
}

func (input *DirectMessageInput) Decode(message []byte) (*DirectMessageInput, error) {
	err := json.Unmarshal(message, input)
	return input, err
}

type DirectMessageOutput struct {
	Sender   *Member         `json:"sender"`
	Receiver *Member         `json:"receiver"`
	Content  json.RawMessage `json:"content"`
}

func (output *DirectMessageOutput) Encode() ([]byte, error) {
	return json.Marshal(output)
}

func NewDirectMessageOutput(sender, receiver *Member, content json.RawMessage) *DirectMessageOutput {
	return &DirectMessageOutput{
		Sender:   sender,
		Receiver: receiver,
		Content:  content,
	}
}
