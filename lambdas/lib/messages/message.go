package messages

import "encoding/json"

type Message struct {
	Sender   string          `json:"sender"`
	Receiver string          `json:"receiver"`
	Content  json.RawMessage `json:"content"`
}

func (message *Message) Encode() ([]byte, error) {
	return json.Marshal(message)
}
