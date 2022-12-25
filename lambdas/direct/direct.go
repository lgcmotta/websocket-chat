package main

import (
	"context"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/lgcmotta/websocket-chat/lib/apigw"
	"github.com/lgcmotta/websocket-chat/lib/config"
	"github.com/lgcmotta/websocket-chat/lib/db"
	"github.com/lgcmotta/websocket-chat/lib/logger"
	"github.com/lgcmotta/websocket-chat/lib/messages"
)

func main() {
	lambda.Start(handle)
}

func handle(ctx context.Context, req *events.APIGatewayWebsocketProxyRequest) (apigw.Response, error) {
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

	receivers := make([]*db.Member, 0)

	for _, conncetionId := range directInput.Receivers {

		receiver, err := db.Instance.GetMember(ctx, conncetionId)

		if err != nil {
			logger.Log.FailedToRetrieveReceiver(req, err)

			return apigw.InternalServerErrorResponse(), err
		}
		receivers = append(receivers, receiver)

	}

	sender, err := db.Instance.GetMember(ctx, req.RequestContext.ConnectionID)

	if err != nil {
		logger.Log.FailedToRetrieveSender(req, err)

		return apigw.InternalServerErrorResponse(), err
	}

	receivedAt := time.Now()

	replyToSender := messages.NewMessageOutput(
		sender.Cast(),
		sender.Cast(),
		directInput.Content,
		&receivedAt,
		"direct",
	)

	apigw.Client.SendMessage(ctx, replyToSender)

	for _, receiver := range receivers {

		direct := messages.NewMessageOutput(
			sender.Cast(),
			receiver.Cast(),
			directInput.Content,
			&receivedAt,
			"direct",
		)

		apigw.Client.SendMessage(ctx, direct)
	}

	return apigw.OkResponse(), nil
}
