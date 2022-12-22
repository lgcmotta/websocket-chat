package messages

import "encoding/json"

type MemberJoinInput struct {
	Nickname string `json:"nickname"`
}

func (input *MemberJoinInput) Decode(message []byte) (*MemberJoinInput, error) {
	err := json.Unmarshal(message, input)
	return input, err
}

type BroadcastMessageInput struct {
	Content json.RawMessage `json:"content"`
}

func (input *BroadcastMessageInput) Decode(message []byte) (*BroadcastMessageInput, error) {
	err := json.Unmarshal(message, input)
	return input, err
}

type DirectMessageInput struct {
	Receiver string          `json:"receiver"`
	Content  json.RawMessage `json:"content"`
}

func (input *DirectMessageInput) Decode(message []byte) (*DirectMessageInput, error) {
	err := json.Unmarshal(message, input)
	return input, err
}

type Message struct {
	Sender   string          `json:"sender"`
	Receiver string          `json:"receiver"`
	Content  json.RawMessage `json:"content"`
}

func (message *Message) Encode() ([]byte, error) {
	return json.Marshal(message)
}
