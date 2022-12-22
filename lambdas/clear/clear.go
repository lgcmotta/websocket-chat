package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/lgcmotta/websocket-chat/lib/apigw"
	"github.com/lgcmotta/websocket-chat/lib/db"
	"github.com/lgcmotta/websocket-chat/lib/logger"

	"go.uber.org/zap"
)

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, req *events.APIGatewayWebsocketProxyRequest) (apigw.Response, error) {
	defer func() {
		_ = logger.Instance.Sync()
	}()

	logger.Instance.Info("websocket disconnecting inactive connections",
		zap.String("requestId", req.RequestContext.RequestID),
		zap.String("connectionId", req.RequestContext.ConnectionID))

	connectionIds, err := db.Instance.GetConnectionIDs(ctx)

	if err != nil {
		logger.Instance.Error("failed to retrieve connection ids",
			zap.String("requestId", req.RequestContext.RequestID),
			zap.String("connectionId", req.RequestContext.ConnectionID),
			zap.Error(err),
		)
	}

	logger.Instance.Info("websocket connections read from cache",
		zap.Int("connections", len(connectionIds)),
		zap.String("requestId", req.RequestContext.RequestID),
		zap.String("connectionId", req.RequestContext.ConnectionID),
	)

	for _, connectionId := range connectionIds {
		if connectionId.ConnectionId == req.RequestContext.ConnectionID {
			continue
		}

		err := db.Instance.RemoveConnectionID(ctx, connectionId.ConnectionId)

		if err != nil {
			logger.Instance.Error("failed to delete connection details from cache",
				zap.String("requestId", req.RequestContext.RequestID),
				zap.String("connectionId", req.RequestContext.ConnectionID),
				zap.Error(err))

			return apigw.InternalServerErrorResponse(), err
		}
	}

	logger.Instance.Info("websocket connections deleted from cache",
		zap.String("requestId", req.RequestContext.RequestID),
		zap.String("connectionId", req.RequestContext.ConnectionID))

	return apigw.OkResponse(), nil
}
