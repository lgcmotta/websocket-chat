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
	lambda.Start(handle)
}

func handle(ctx context.Context, req *events.APIGatewayWebsocketProxyRequest) (apigw.Response, error) {
	defer logger.Log.RequestEnded(req)

	logger.Log.RequestStarted(req)

	if apigw.Client == nil {
		apigw.Client = apigw.NewAPIGatewayManagementClient(&config.Configuration, req.RequestContext.DomainName, req.RequestContext.Stage)
	}

	member, err := db.Instance.GetMember(ctx, req.RequestContext.ConnectionID)

	if err != nil {
		logger.Log.FailedToRetrieveLeavingMember(req, err)

		return apigw.InternalServerErrorResponse(), err
	}

	err = db.Instance.RemoveConnectionID(ctx, req.RequestContext.ConnectionID)

	if err != nil {
		logger.Log.FailedToDeleteConnectionId(req, err)

		return apigw.InternalServerErrorResponse(), err
	}

	connectedMembers, err := db.Instance.GetMembers(ctx)

	if err != nil {
		logger.Log.FailedToRetrieveMembers(req, err)

		return apigw.InternalServerErrorResponse(), err
	}

	receivers := make([]*messages.Member, 0)

	for _, connectedMember := range connectedMembers {
		receivers = append(receivers, connectedMember.Cast())
	}

	sender := member.Cast()

	broadcast := messages.NewBroadcastMessageOutput(sender, receivers, sender.GetLeavingMessage())

	apigw.Client.BroadcastMessage(ctx, broadcast)

	return apigw.OkResponse(), nil
}
