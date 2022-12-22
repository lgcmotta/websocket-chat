package db

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/lgcmotta/websocket-chat/lib/config"
	"github.com/lgcmotta/websocket-chat/lib/logger"
)

type DbClient struct {
	dynamoDbClient    *dynamodb.Client
	dynamoDBTableName *string
}

var Instance *DbClient

func init() {
	Instance = &DbClient{
		dynamoDbClient:    dynamodb.NewFromConfig(config.Configuration),
		dynamoDBTableName: aws.String(os.Getenv("DYNAMODB_TABLE_NAME")),
	}
}

func (client *DbClient) AddConnectionID(ctx context.Context, connectionID string) error {
	connection := Member{ConnectionId: connectionID}

	item, err := attributevalue.MarshalMap(connection)

	if err != nil {
		logger.Log.FailedToMarshalItemDynamoDB(err)
	}

	input := &dynamodb.PutItemInput{
		TableName: client.dynamoDBTableName,
		Item:      item,
	}

	_, err = client.dynamoDbClient.PutItem(ctx, input)

	if err != nil {
		logger.Log.FailedToPutItemDynamoDB(connectionID, err)
	}

	return err
}

func (client *DbClient) GetMembers(ctx context.Context) ([]Member, error) {
	input := &dynamodb.ScanInput{
		TableName: client.dynamoDBTableName,
	}

	var members []Member

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

func (client *DbClient) GetMember(ctx context.Context, connectionID string) (*Member, error) {
	member := new(Member)
	member.ConnectionId = connectionID

	input := &dynamodb.GetItemInput{
		Key:       member.GetKey(),
		TableName: client.dynamoDBTableName,
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
	connection := Member{ConnectionId: connectionID}

	input := &dynamodb.DeleteItemInput{
		TableName: client.dynamoDBTableName, Key: connection.GetKey(),
	}
	_, err := client.dynamoDbClient.DeleteItem(ctx, input)

	if err != nil {
		logger.Log.FailedToDeleteItemDynamoDB(connectionID, err)
	}

	return err
}

func (client *DbClient) SetMemberName(ctx context.Context, connectionID, name string) error {
	member := Member{ConnectionId: connectionID, Nickname: name}

	query := expression.Set(expression.Name("nickname"), expression.Value(member.Nickname))

	expr, err := expression.NewBuilder().WithUpdate(query).Build()

	if err != nil {
		logger.Log.FailedToPutItemDynamoDB(connectionID, err)

		return err
	}

	input := &dynamodb.UpdateItemInput{
		TableName:                 client.dynamoDBTableName,
		Key:                       member.GetKey(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
		ReturnValues:              types.ReturnValueUpdatedNew,
	}

	client.dynamoDbClient.UpdateItem(ctx, input)
	return nil
}
