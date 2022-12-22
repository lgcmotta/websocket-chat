package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/lgcmotta/websocket-chat/lib/apigw"
	"github.com/lgcmotta/websocket-chat/lib/config"
	"github.com/lgcmotta/websocket-chat/lib/db"
	"github.com/lgcmotta/websocket-chat/lib/logger"
	"github.com/lgcmotta/websocket-chat/lib/members"
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

	logger.Instance.Info("websocket join",
		zap.String("requestId", req.RequestContext.RequestID),
		zap.String("connectionId", req.RequestContext.ConnectionID),
	)

	joinInput, err := new(messages.MemberJoinInput).Decode([]byte(req.Body))

	if err != nil {
		logger.Instance.Error("failed to decode client input",
			zap.String("requestId", req.RequestContext.RequestID),
			zap.String("connectionId", req.RequestContext.ConnectionID),
			zap.Error(err),
		)

		return apigw.BadRequestResponse(), err
	}

	err = db.Instance.SetMemberName(ctx, req.RequestContext.ConnectionID, joinInput.Nickname)

	if err != nil {
		logger.Instance.Error("failed to update client nickname",
			zap.String("requestId", req.RequestContext.RequestID),
			zap.String("connectionId", req.RequestContext.ConnectionID),
			zap.String("nickname", joinInput.Nickname),
			zap.Error(err),
		)

		return apigw.InternalServerErrorResponse(), err
	}

	connectedMembers, err := db.Instance.GetMembers(ctx)

	if err != nil {
		logger.Instance.Error("failed to retrive chat members",
			zap.String("requestId", req.RequestContext.RequestID),
			zap.String("connectionId", req.RequestContext.ConnectionID),
			zap.String("nickname", joinInput.Nickname),
			zap.Error(err),
		)

		return apigw.InternalServerErrorResponse(), err
	}

	sender := members.Member{
		ConnectionId: req.RequestContext.ConnectionID,
		Nickname:     joinInput.Nickname,
	}

	message := fmt.Sprintf("%s joined the chat", joinInput.Nickname)

	output := messages.NewMemberJoinOutput(sender.ConnectionId, sender.Nickname, message)

	encodedOutput, err := output.Encode()

	if err != nil {
		logger.Instance.Error("failed to encode client output",
			zap.String("requestId", req.RequestContext.RequestID),
			zap.String("connectionId", req.RequestContext.ConnectionID),
			zap.Error(err),
		)

		return apigw.BadRequestResponse(), err
	}

	apigw.Client.BroadcastMessage(ctx, sender, connectedMembers, encodedOutput)

	return apigw.OkResponse(), nil
}
