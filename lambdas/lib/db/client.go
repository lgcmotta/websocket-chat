package db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/lgcmotta/websocket-chat/lib/logger"
	"github.com/lgcmotta/websocket-chat/lib/members"
	"go.uber.org/zap"
)

type DbClient struct {
	dynamoDbClient *dynamodb.Client
}

var Instance *DbClient

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		logger.Instance.Error("failed to get default config for dynamodb client",
			zap.Error(err),
		)
	}

	Instance = &DbClient{
		dynamoDbClient: dynamodb.NewFromConfig(cfg),
	}
}

func (client *DbClient) AddConnectionID(ctx context.Context, connectionID string) error {
	connection := members.Member{ConnectionId: connectionID}

	item, err := attributevalue.MarshalMap(connection)

	if err != nil {
		logger.Instance.Error("failed to marshal connection id as a map",
			zap.Error(err),
		)
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String("ConnectionIds"),
		Item:      item,
	}

	_, err = client.dynamoDbClient.PutItem(ctx, input)

	if err != nil {
		logger.Instance.Error("failed to connection id on dynamodb",
			zap.Error(err),
		)
	}

	return err
}

func (client *DbClient) GetMembers(ctx context.Context) ([]members.Member, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String("ConnectionIds"),
	}

	var members []members.Member

	result, err := client.dynamoDbClient.Scan(ctx, input)

	if err != nil {
		logger.Instance.Error("failed to scan connection ids on dynamodb",
			zap.Error(err),
		)
	}

	err = attributevalue.UnmarshalListOfMaps(result.Items, &members)

	if err != nil {
		logger.Instance.Error("failed to unmarshal connection ids from dynamodb",
			zap.Error(err),
		)
	}

	return members, nil
}

func (client *DbClient) GetMember(ctx context.Context, connectionID string) (*members.Member, error) {
	member := new(members.Member)
	member.ConnectionId = connectionID

	input := &dynamodb.GetItemInput{
		Key:       member.GetKey(),
		TableName: aws.String("ConnectionIds"),
	}

	response, err := client.dynamoDbClient.GetItem(ctx, input)

	if err != nil {
		logger.Instance.Error("failed to get member",
			zap.String("connectionId", connectionID),
			zap.Error(err),
		)
		return nil, err
	}

	err = attributevalue.UnmarshalMap(response.Item, member)

	if err != nil {
		logger.Instance.Error("unmarshal member failed",
			zap.String("connectionId", connectionID),
			zap.Error(err),
		)
		return nil, err
	}

	return member, nil
}

func (client *DbClient) RemoveConnectionID(ctx context.Context, connectionID string) error {
	connection := members.Member{ConnectionId: connectionID}

	input := &dynamodb.DeleteItemInput{
		TableName: aws.String("ConnectionIds"), Key: connection.GetKey(),
	}
	_, err := client.dynamoDbClient.DeleteItem(ctx, input)

	if err != nil {
		logger.Instance.Error("failed to remove connection id from dynamodb",
			zap.Error(err),
		)
	}

	return err
}

func (client *DbClient) SetMemberName(ctx context.Context, connectionID, name string) error {
	member := members.Member{ConnectionId: connectionID, Nickname: name}

	query := expression.Set(expression.Name("nickname"), expression.Value(member.Nickname))

	expr, err := expression.NewBuilder().WithUpdate(query).Build()

	if err != nil {
		logger.Instance.Error("build set member name update query failed",
			zap.String("connectionId", connectionID),
			zap.String("name", name),
			zap.Error(err),
		)
		return err
	}

	input := &dynamodb.UpdateItemInput{
		TableName:                 aws.String("ConnectionIds"),
		Key:                       member.GetKey(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
		ReturnValues:              types.ReturnValueUpdatedNew,
	}

	client.dynamoDbClient.UpdateItem(ctx, input)
	return nil
}
