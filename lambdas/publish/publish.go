package main

import (
	"context"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/mediocregopher/radix/v4"
	"go.uber.org/zap"

	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/lgcmotta/websocket-chat/lib/apigw"
	"github.com/lgcmotta/websocket-chat/lib/apigw/ws"
	"github.com/lgcmotta/websocket-chat/lib/connections"
	"github.com/lgcmotta/websocket-chat/lib/logger"
	"github.com/lgcmotta/websocket-chat/lib/redis"
)

var cfg aws.Config

var apiClient *apigatewaymanagementapi.Client

func init() {
	var err error
	cfg, err = config.LoadDefaultConfig(context.TODO())
	if err != nil {
		logger.Instance.Panic("unable to load SDK config", zap.Error(err))
	}
}

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, req *events.APIGatewayWebsocketProxyRequest) (apigw.Response, error) {
	defer func() {
		_ = logger.Instance.Sync()
	}()

	childContext, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	logger.Instance.Info("websocket publish",
		zap.String("requestId", req.RequestContext.RequestID),
		zap.String("connectionId", req.RequestContext.ConnectionID),
		zap.String("stage", req.RequestContext.Stage),
		zap.String("domain-name", req.RequestContext.DomainName),
		zap.String("account-id", req.RequestContext.AccountID),
		zap.String("api-id", req.RequestContext.APIID),
	)

	if apiClient == nil {
		apiClient = apigw.NewAPIGatewayManagementClient(&cfg, req.RequestContext.DomainName, req.RequestContext.Stage)
	}

	input, err := new(ws.InputEnvelop).Decode([]byte(req.Body))

	logger.Instance.Info("input envelope",
		zap.Bool("echo", input.Echo),
		zap.String("data", string(input.Data)),
		zap.Int("type", input.Type),
	)

	if err != nil {
		logger.Instance.Error("failed to parse client input",
			zap.String("requestId", req.RequestContext.RequestID),
			zap.String("connectionId", req.RequestContext.ConnectionID),
			zap.Error(err))

		return apigw.BadRequestResponse(), err
	}

	output := &ws.OutputEnvelop{
		Data:     input.Data,
		Type:     input.Type,
		Received: time.Now().Unix(),
	}

	data, err := output.Encode()
	if err != nil {
		logger.Instance.Error("failed to encode output",
			zap.String("requestId", req.RequestContext.RequestID),
			zap.String("connectionId", req.RequestContext.ConnectionID),
			zap.Error(err))

		return apigw.InternalServerErrorResponse(), err
	}

	connected := new(connections.Connections)
	err = redis.Client.Do(childContext, radix.Cmd(&(connected.ConnectionIds), "SMEMBERS", "connections"))
	if err != nil {
		logger.Instance.Error("failed to read connections from cache",
			zap.String("requestId", req.RequestContext.RequestID),
			zap.String("connectionId", req.RequestContext.ConnectionID),
			zap.Error(err))

		return apigw.InternalServerErrorResponse(), err
	}

	logger.Instance.Info("websocket connections read from cache",
		zap.Int("connections", connected.Len()),
		zap.String("requestId", req.RequestContext.RequestID),
		zap.String("connectionId", req.RequestContext.ConnectionID))

	for _, connectionId := range connected.ConnectionIds {
		err = handleError(childContext, publish(childContext, connectionId, data), connectionId)

		if err != nil {
			logger.Instance.Error("failed to publish to connection",
				zap.String("receiver", connectionId),
				zap.String("requestId", req.RequestContext.RequestID),
				zap.String("sender", req.RequestContext.ConnectionID),
				zap.Error(err))
		}
	}

	return apigw.OkResponse(), nil
}

func publish(ctx context.Context, connectionId string, data []byte) error {
	_, err := apiClient.PostToConnection(ctx, &apigatewaymanagementapi.PostToConnectionInput{
		Data:         data,
		ConnectionId: aws.String(connectionId),
	})

	return err
}

func handleError(ctx context.Context, err error, connectionId string) error {
	if err == nil {
		return err
	}

	return deleteConnectionId(ctx, connectionId)
}

func deleteConnectionId(ctx context.Context, id string) error {
	var result string
	err := redis.Client.Do(ctx, radix.Cmd(&result, "SREM", "connections", id))
	if err != nil {
		logger.Instance.Error("failed to delete connection details from cache",
			zap.String("connectionId", id),
			zap.Error(err))

		return err
	}

	logger.Instance.Info("websocket connection deleted from cache",
		zap.String("result", result),
		zap.String("connectionId", id))

	return err
}
