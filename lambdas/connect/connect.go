package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/lgcmotta/websocket-chat/lib/apigw"
	"github.com/lgcmotta/websocket-chat/lib/logger"
	"github.com/lgcmotta/websocket-chat/lib/redis"
	"github.com/mediocregopher/radix/v4"
	"go.uber.org/zap"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, req *events.APIGatewayWebsocketProxyRequest) (apigw.Response, error) {
	defer func() {
		_ = logger.Instance.Sync()
	}()

	logger.Instance.Info("websocket connect",
		zap.String("requestId", req.RequestContext.RequestID),
		zap.String("connectionId", req.RequestContext.ConnectionID),
	)

	var result string
	err := redis.Client.Do(ctx, radix.Cmd(&result, "SADD", "connections", req.RequestContext.ConnectionID))

	if err != nil {
		logger.Instance.Error("failed to cache connection details",
			zap.String("requestId", req.RequestContext.RequestID),
			zap.String("connectionId", req.RequestContext.ConnectionID),
			zap.Error(err))

		return apigw.InternalServerErrorResponse(), err
	}

	logger.Instance.Info("websocket connection cached",
		zap.String("result", result),
		zap.String("requestId", req.RequestContext.RequestID),
		zap.String("connectionId", req.RequestContext.ConnectionID))

	return apigw.OkResponse(), nil
}
