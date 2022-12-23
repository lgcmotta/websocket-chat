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

	joinInput, err := new(messages.MemberJoinInput).Decode([]byte(req.Body))

	if err != nil {
		logger.Log.FailedToDecodeInput(req, err)

		return apigw.BadRequestResponse(), err
	}

	err = db.Instance.SetMemberName(ctx, req.RequestContext.ConnectionID, joinInput.Nickname)

	if err != nil {
		logger.Log.FailedToUpdateMemberNickname(req, joinInput.Nickname, err)

		return apigw.InternalServerErrorResponse(), err
	}

	connectedMembers, err := db.Instance.GetMembers(ctx)

	if err != nil {
		logger.Log.FailedToRetrieveMembers(req, err)

		return apigw.InternalServerErrorResponse(), err
	}

	sender := messages.NewMember(req.RequestContext.ConnectionID, joinInput.Nickname)

	receivers := make([]*messages.Member, 0)

	for _, connectedMember := range connectedMembers {
		receivers = append(receivers, connectedMember.Cast())
	}

	receivedAt := time.Now()

	broadcast := messages.NewBroadcastMessageOutput(sender, receivers, sender.GetJoiningMessage(), &receivedAt)

	apigw.Client.BroadcastMessage(ctx, broadcast)

	return apigw.OkResponse(), nil
}
