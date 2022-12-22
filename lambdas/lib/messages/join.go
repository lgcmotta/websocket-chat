package messages

import "encoding/json"

type MemberJoinInput struct {
	Nickname string `json:"nickname"`
}

func (input *MemberJoinInput) Decode(message []byte) (*MemberJoinInput, error) {
	err := json.Unmarshal(message, input)
	return input, err
}
