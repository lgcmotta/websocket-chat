package connections

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type ConnectionId struct {
	ConnectionId string `dynamodbav:"connectionId"`
}

func (connection ConnectionId) GetKey() map[string]types.AttributeValue {
	connectionId, err := attributevalue.Marshal(connection.ConnectionId)
	if err != nil {
		panic(err)
	}
	return map[string]types.AttributeValue{"connectionId": connectionId}
}
