package db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/lgcmotta/websocket-chat/lib/config"
	"github.com/lgcmotta/websocket-chat/lib/logger"
	"github.com/lgcmotta/websocket-chat/lib/members"
)

type DbClient struct {
	dynamoDbClient *dynamodb.Client
}

var Instance *DbClient

func init() {
	Instance = &DbClient{
		dynamoDbClient: dynamodb.NewFromConfig(config.Configuration),
	}
}

func (client *DbClient) AddConnectionID(ctx context.Context, connectionID string) error {
	connection := members.Member{ConnectionId: connectionID}

	item, err := attributevalue.MarshalMap(connection)

	if err != nil {
		logger.Log.FailedToMarshalItemDynamoDB(err)
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String("ConnectionIds"),
		Item:      item,
	}

	_, err = client.dynamoDbClient.PutItem(ctx, input)

	if err != nil {
		logger.Log.FailedToPutItemDynamoDB(connectionID, err)
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
		logger.Log.FailedToScanItemsDynamoDB(err)
	}

	err = attributevalue.UnmarshalListOfMaps(result.Items, &members)

	if err != nil {
		logger.Log.FailedToUnmarshalItemDynamoDB(err)
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
		logger.Log.FailedToGetItemDynamoDB(connectionID, err)
		return nil, err
	}

	err = attributevalue.UnmarshalMap(response.Item, member)

	if err != nil {
		logger.Log.FailedToUnmarshalItemDynamoDB(err)
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
		logger.Log.FailedToDeleteItemDynamoDB(connectionID, err)
	}

	return err
}

func (client *DbClient) SetMemberName(ctx context.Context, connectionID, name string) error {
	member := members.Member{ConnectionId: connectionID, Nickname: name}

	query := expression.Set(expression.Name("nickname"), expression.Value(member.Nickname))

	expr, err := expression.NewBuilder().WithUpdate(query).Build()

	if err != nil {
		logger.Log.FailedToPutItemDynamoDB(connectionID, err)

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
