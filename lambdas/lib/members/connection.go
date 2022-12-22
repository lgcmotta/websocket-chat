package members

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Member struct {
	ConnectionId string `dynamodbav:"connectionId"`
}

func (member Member) GetKey() map[string]types.AttributeValue {
	connectionId, err := attributevalue.Marshal(member.ConnectionId)
	if err != nil {
		panic(err)
	}
	return map[string]types.AttributeValue{"connectionId": connectionId}
}
