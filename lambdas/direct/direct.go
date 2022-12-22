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

	directInput, err := new(messages.DirectMessageInput).Decode([]byte(req.Body))

	if err != nil {
		logger.Instance.Error("failed to decode client input",
			zap.String("requestId", req.RequestContext.RequestID),
			zap.String("connectionId", req.RequestContext.ConnectionID),
			zap.Error(err),
		)

		return apigw.BadRequestResponse(), err
	}

	receiver, err := db.Instance.GetMember(ctx, directInput.Receiver)

	if err != nil {
		logger.Instance.Error("failed to retrieve receiver",
			zap.String("requestId", req.RequestContext.RequestID),
			zap.String("connectionId", req.RequestContext.ConnectionID),
			zap.Error(err),
		)

		return apigw.InternalServerErrorResponse(), err
	}

	sender, err := db.Instance.GetMember(ctx, req.RequestContext.ConnectionID)

	if err != nil {
		logger.Instance.Error("failed to retrieve sender",
			zap.String("requestId", req.RequestContext.RequestID),
			zap.String("connectionId", req.RequestContext.ConnectionID),
			zap.Error(err),
		)

		return apigw.InternalServerErrorResponse(), err
	}

	direct := messages.NewDirectMessageOutput(
		sender.Cast(),
		receiver.Cast(),
		directInput.Content,
	)

	apigw.Client.SendPrivateMessage(ctx, direct)

	return apigw.OkResponse(), nil
}
