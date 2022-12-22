package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/lgcmotta/websocket-chat/lib/apigw"
	"github.com/lgcmotta/websocket-chat/lib/db"
	"github.com/lgcmotta/websocket-chat/lib/logger"
)

func main() {
	lambda.Start(handle)
}

func handle(ctx context.Context, req *events.APIGatewayWebsocketProxyRequest) (apigw.Response, error) {
	defer logger.Log.RequestEnded(req)

	logger.Log.RequestStarted(req)

	err := db.Instance.AddConnectionID(ctx, req.RequestContext.ConnectionID)

	if err != nil {
		logger.Log.FailedToSaveConnectionId(req, err)

		return apigw.InternalServerErrorResponse(), err
	}

	logger.Log.ConnectionIdSaved(req)

	return apigw.OkResponse(), nil
}
