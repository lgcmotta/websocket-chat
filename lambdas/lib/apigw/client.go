package apigw

import (
	"context"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/lgcmotta/websocket-chat/lib/logger"
	"github.com/lgcmotta/websocket-chat/lib/members"
	"github.com/lgcmotta/websocket-chat/lib/messages"
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

type ApiClient struct {
	client *apigatewaymanagementapi.Client
}

var Client *ApiClient

func NewAPIGatewayManagementClient(cfg *aws.Config, domain, stage string) *ApiClient {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		var endpoint url.URL
		endpoint.Path = stage
		endpoint.Host = domain
		endpoint.Scheme = "https"

		return aws.Endpoint{
			SigningRegion: region,
			URL:           endpoint.String(),
			PartitionID:   "aws",
		}, nil
	})

	cfg.EndpointResolverWithOptions = customResolver

	client := apigatewaymanagementapi.NewFromConfig(*cfg)

	return &ApiClient{
		client: client,
	}
}

func (client *ApiClient) BroadcastMessage(ctx context.Context, sender members.Member, receivers []members.Member, content []byte) error {
	var errs error
	for _, receiver := range receivers {
		err := client.SendPrivateMessage(ctx, sender, receiver, content)
		if err != nil {
			errs = multierr.Append(errs, err)
		}
	}
	return errs
}

func (client *ApiClient) SendPrivateMessage(ctx context.Context, sender members.Member, receiver members.Member, content []byte) error {
	var errs error
	message := messages.Message{
		Sender:   sender.Nickname,
		Receiver: receiver.Nickname,
		Content:  content,
	}

	encoded, err := message.Encode()

	if err != nil {
		logger.Instance.Error("encode message failed",
			zap.String("receiverId", receiver.ConnectionId),
			zap.String("receiverName", receiver.Nickname),
			zap.String("senderId", sender.ConnectionId),
			zap.String("senderName", sender.Nickname),
			zap.Error(err),
		)
		errs = multierr.Append(errs, err)
	}

	input := &apigatewaymanagementapi.PostToConnectionInput{
		ConnectionId: aws.String(receiver.ConnectionId),
		Data:         encoded,
	}

	_, err = client.client.PostToConnection(ctx, input)

	if err != nil {
		logger.Instance.Error("send message failed",
			zap.String("receiverId", receiver.ConnectionId),
			zap.String("receiverName", receiver.Nickname),
			zap.String("senderId", sender.ConnectionId),
			zap.String("senderName", sender.Nickname),
			zap.Error(err),
		)
		errs = multierr.Append(errs, err)
	}

	return errs
}
