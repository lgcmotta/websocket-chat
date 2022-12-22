package db

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/lgcmotta/websocket-chat/lib/messages"
)

type Member struct {
	ConnectionId string `dynamodbav:"connectionId"`
	Nickname     string `dynamodbav:"nickname"`
}

func (member Member) GetKey() map[string]types.AttributeValue {
	connectionId, err := attributevalue.Marshal(member.ConnectionId)
	if err != nil {
		panic(err)
	}
	return map[string]types.AttributeValue{"connectionId": connectionId}
}

func (member *Member) Cast() *messages.Member {
	return messages.NewMember(
		member.ConnectionId,
		member.Nickname,
	)
}
