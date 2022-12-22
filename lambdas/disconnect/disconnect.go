package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/lgcmotta/websocket-chat/lib/apigw"
	"github.com/lgcmotta/websocket-chat/lib/db"
	"github.com/lgcmotta/websocket-chat/lib/logger"
	"go.uber.org/zap"
)

var cfg aws.Config

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
	defer logger.Sync()

	if apigw.Client == nil {
		apigw.Client = apigw.NewAPIGatewayManagementClient(&cfg, req.RequestContext.DomainName, req.RequestContext.Stage)
	}

	logger.Instance.Info("websocket disconnect",
		zap.String("requestId", req.RequestContext.RequestID),
		zap.String("connectionId", req.RequestContext.ConnectionID))

	member, err := db.Instance.GetMember(ctx, req.RequestContext.ConnectionID)

	if err != nil {
		logger.Instance.Error("failed to retrieve leaving member",
			zap.String("requestId", req.RequestContext.RequestID),
			zap.String("connectionId", req.RequestContext.ConnectionID),
			zap.Error(err),
		)

		return apigw.InternalServerErrorResponse(), err
	}

	err = db.Instance.RemoveConnectionID(ctx, req.RequestContext.ConnectionID)

	if err != nil {
		logger.Instance.Error("failed to delete connection details",
			zap.String("requestId", req.RequestContext.RequestID),
			zap.String("connectionId", req.RequestContext.ConnectionID),
			zap.Error(err),
		)

		return apigw.InternalServerErrorResponse(), err
	}

	connectedMembers, err := db.Instance.GetMembers(ctx)

	if err != nil {
		logger.Instance.Error("failed to get connected members",
			zap.String("requestId", req.RequestContext.RequestID),
			zap.String("connectionId", req.RequestContext.ConnectionID),
			zap.Error(err))

		return apigw.InternalServerErrorResponse(), err
	}

	message := fmt.Sprintf("%s left the chat", member.Nickname)

	apigw.Client.BroadcastMessage(ctx, *member, connectedMembers, []byte(message))

	logger.Instance.Info("websocket connection deleted",
		zap.String("requestId", req.RequestContext.RequestID),
		zap.String("connectionId", req.RequestContext.ConnectionID))

	return apigw.OkResponse(), nil
}
