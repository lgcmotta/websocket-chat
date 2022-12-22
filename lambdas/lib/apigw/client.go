package apigw

import (
	"context"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/lgcmotta/websocket-chat/lib/logger"
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

func (client *ApiClient) BroadcastMessage(ctx context.Context, message *messages.BroadcastMessageOutput) error {
	var errs error
	for _, receiver := range message.Receivers {
		direct := messages.NewDirectMessageOutput(message.Sender, receiver, message.Content)

		err := client.SendPrivateMessage(ctx, direct)

		if err != nil {
			errs = multierr.Append(errs, err)
		}
	}
	return errs
}

func (client *ApiClient) SendPrivateMessage(ctx context.Context, message *messages.DirectMessageOutput) error {
	var errs error

	encoded, err := message.Encode()

	if err != nil {
		logger.Instance.Error("encode message failed",
			zap.String("receiverId", message.Receiver.ConnectionId),
			zap.String("receiverName", message.Receiver.Nickname),
			zap.String("senderId", message.Sender.ConnectionId),
			zap.String("senderName", message.Sender.Nickname),
			zap.Error(err),
		)
		errs = multierr.Append(errs, err)
	}

	input := &apigatewaymanagementapi.PostToConnectionInput{
		ConnectionId: aws.String(message.Receiver.ConnectionId),
		Data:         encoded,
	}

	_, err = client.client.PostToConnection(ctx, input)

	if err != nil {
		logger.Instance.Error("send message failed",
			zap.String("receiverId", message.Receiver.ConnectionId),
			zap.String("receiverName", message.Receiver.Nickname),
			zap.String("senderId", message.Sender.ConnectionId),
			zap.String("senderName", message.Sender.Nickname),
			zap.Error(err),
		)
		errs = multierr.Append(errs, err)
	}

	return errs
}
