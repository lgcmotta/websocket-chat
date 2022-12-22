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
)

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, req *events.APIGatewayWebsocketProxyRequest) (apigw.Response, error) {
	defer logger.Log.RequestEnded(req)

	logger.Log.RequestStarted(req)

	if apigw.Client == nil {
		apigw.Client = apigw.NewAPIGatewayManagementClient(&config.Configuration, req.RequestContext.DomainName, req.RequestContext.Stage)
	}

	directInput, err := new(messages.DirectMessageInput).Decode([]byte(req.Body))

	if err != nil {
		logger.Log.FailedToDecodeInput(req, err)

		return apigw.BadRequestResponse(), err
	}

	receiver, err := db.Instance.GetMember(ctx, directInput.Receiver)

	if err != nil {
		logger.Log.FailedToRetrieveReceiver(req, err)

		return apigw.InternalServerErrorResponse(), err
	}

	sender, err := db.Instance.GetMember(ctx, req.RequestContext.ConnectionID)

	if err != nil {
		logger.Log.FailedToRetrieveSender(req, err)

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
