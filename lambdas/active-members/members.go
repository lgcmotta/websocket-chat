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

	connectedMembers, err := db.Instance.GetMembers(ctx)

	if err != nil {
		logger.Log.FailedToRetrieveMembers(req, err)

		return apigw.InternalServerErrorResponse(), err
	}

	member, err := db.Instance.GetMember(ctx, req.RequestContext.ConnectionID)

	if err != nil {
		logger.Log.FailedToRetrieveSender(req, err)

		return apigw.InternalServerErrorResponse(), err
	}

	receiver := messages.NewMember(member.ConnectionId, member.Nickname)

	members := new(messages.ConnectedMembers)

	members.Members = make([]*messages.Member, 0)

	for _, connectedMember := range connectedMembers {
		members.Members = append(members.Members, connectedMember.Cast())
	}

	apigw.Client.SendConnectedClients(ctx, receiver, members)

	return apigw.OkResponse(), nil
}
