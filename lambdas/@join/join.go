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
	"github.com/lgcmotta/websocket-chat/lib/members"
	"github.com/lgcmotta/websocket-chat/lib/messages"
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

	logger.Instance.Info("websocket join",
		zap.String("requestId", req.RequestContext.RequestID),
		zap.String("connectionId", req.RequestContext.ConnectionID),
	)

	joinInput, err := new(messages.MemberJoinInput).Decode([]byte(req.Body))

	if err != nil {
		logger.Instance.Error("failed to parse client input",
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

	apigw.Client.BroadcastMessage(ctx, sender, connectedMembers, []byte(message))

	return apigw.OkResponse(), nil
}
