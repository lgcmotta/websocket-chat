package messages

import "encoding/json"

type MemberJoinInput struct {
	Nickname string `json:"nickname"`
}

func (input *MemberJoinInput) Decode(message []byte) (*MemberJoinInput, error) {
	err := json.Unmarshal(message, input)
	return input, err
}

type MemberJoinOutput struct {
	ConnectionId string `json:"connectionId"`
	Nickname     string `json:"nickname"`
	Message      string `json:"message"`
}

func NewMemberJoinOutput(connectionId, nickname, message string) *MemberJoinOutput {
	return &MemberJoinOutput{
		ConnectionId: connectionId,
		Nickname:     nickname,
		Message:      message,
	}
}

func (output *MemberJoinOutput) Encode() ([]byte, error) {
	return json.Marshal(output)
}
