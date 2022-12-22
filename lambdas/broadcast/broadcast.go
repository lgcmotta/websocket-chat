package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/lgcmotta/websocket-chat/lib/apigw"
	"github.com/lgcmotta/websocket-chat/lib/config"
	"github.com/lgcmotta/websocket-chat/lib/db"
	"github.com/lgcmotta/websocket-chat/lib/logger"
	"github.com/lgcmotta/websocket-chat/lib/messages"
	"go.uber.org/zap"
)

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, req *events.APIGatewayWebsocketProxyRequest) (apigw.Response, error) {
	defer logger.Sync()

	if apigw.Client == nil {
		apigw.Client = apigw.NewAPIGatewayManagementClient(&config.Configuration, req.RequestContext.DomainName, req.RequestContext.Stage)
	}

	broadcastInput, err := new(messages.BroadcastMessageInput).Decode([]byte(req.Body))

	if err != nil {
		logger.Instance.Error("failed to parse client input",
			zap.String("requestId", req.RequestContext.RequestID),
			zap.String("connectionId", req.RequestContext.ConnectionID),
			zap.Error(err),
		)

		return apigw.BadRequestResponse(), err
	}

	connectedMembers, err := db.Instance.GetMembers(ctx)

	if err != nil {
		logger.Instance.Error("failed to retrive chat members",
			zap.String("requestId", req.RequestContext.RequestID),
			zap.String("connectionId", req.RequestContext.ConnectionID),
			zap.Error(err),
		)

		return apigw.InternalServerErrorResponse(), err
	}

	sender, err := db.Instance.GetMember(ctx, req.RequestContext.ConnectionID)

	if err != nil {
		logger.Instance.Error("failed to retrieve sender member",
			zap.String("requestId", req.RequestContext.RequestID),
			zap.String("connectionId", req.RequestContext.ConnectionID),
			zap.Error(err),
		)

		return apigw.InternalServerErrorResponse(), err
	}

	receivers := make([]*messages.Member, 0)

	for _, connectedMember := range connectedMembers {
		receivers = append(receivers, connectedMember.Cast())
	}

	broadcast := messages.NewBroadcastMessageOutput(sender.Cast(), receivers, broadcastInput.Content)

	apigw.Client.BroadcastMessage(ctx, broadcast)

	return apigw.OkResponse(), nil
}
