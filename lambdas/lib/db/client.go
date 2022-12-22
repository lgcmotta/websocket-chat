package db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/lgcmotta/websocket-chat/lib/connections"
	"github.com/lgcmotta/websocket-chat/lib/logger"
	"go.uber.org/zap"
)

type DbClient struct {
	DynamoDbClient *dynamodb.Client
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
		DynamoDbClient: dynamodb.NewFromConfig(cfg),
	}
}

func (client *DbClient) AddConnectionID(ctx context.Context, connectionID string) error {
	connection := connections.ConnectionId{ConnectionId: connectionID}

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

	_, err = client.DynamoDbClient.PutItem(ctx, input)

	if err != nil {
		logger.Instance.Error("failed to connection id on dynamodb",
			zap.Error(err),
		)
	}

	return err
}

func (client *DbClient) GetConnectionIDs(ctx context.Context) ([]connections.ConnectionId, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String("ConnectionIds"),
	}

	var connectionIds []connections.ConnectionId

	result, err := client.DynamoDbClient.Scan(ctx, input)

	if err != nil {
		logger.Instance.Error("failed to scan connection ids on dynamodb",
			zap.Error(err),
		)
	}

	err = attributevalue.UnmarshalListOfMaps(result.Items, &connectionIds)

	if err != nil {
		logger.Instance.Error("failed to unmarshal connection ids from dynamodb",
			zap.Error(err),
		)
	}

	return connectionIds, nil
}

func (client *DbClient) RemoveConnectionID(ctx context.Context, connectionID string) error {
	connection := connections.ConnectionId{ConnectionId: connectionID}

	input := &dynamodb.DeleteItemInput{
		TableName: aws.String("ConnectionIds"), Key: connection.GetKey(),
	}
	_, err := client.DynamoDbClient.DeleteItem(ctx, input)

	if err != nil {
		logger.Instance.Error("failed to remove connection id from dynamodb",
			zap.Error(err),
		)
	}

	return err
}
