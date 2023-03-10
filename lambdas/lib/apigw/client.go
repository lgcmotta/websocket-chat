package apigw

import (
	"context"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/lgcmotta/websocket-chat/lib/logger"
	"github.com/lgcmotta/websocket-chat/lib/messages"
	"go.uber.org/multierr"
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
		var direct *messages.MessageOutput

		if message.MessageType == "system" {
			system := messages.NewMember("", "@system")
			direct = messages.NewMessageOutput(system, receiver, message.Content, message.ReceivedAt, message.MessageType)
		} else {
			direct = messages.NewMessageOutput(message.Sender, receiver, message.Content, message.ReceivedAt, message.MessageType)
		}

		err := client.SendMessage(ctx, direct)

		if err != nil {
			errs = multierr.Append(errs, err)
		}
	}
	return errs
}

func (client *ApiClient) SendMessage(ctx context.Context, message *messages.MessageOutput) error {
	var errs error

	encoded, err := message.Encode()

	if err != nil {
		logger.Log.FailedToEncodeOutput(
			message.Sender.ConnectionId,
			message.Receiver.ConnectionId,
			err,
		)

		errs = multierr.Append(errs, err)
	}

	input := &apigatewaymanagementapi.PostToConnectionInput{
		ConnectionId: aws.String(message.Receiver.ConnectionId),
		Data:         encoded,
	}

	_, err = client.client.PostToConnection(ctx, input)

	if err != nil {
		logger.Log.FailedToSendMessage(
			message.Sender.ConnectionId,
			message.Receiver.ConnectionId,
			err,
		)
		errs = multierr.Append(errs, err)
	}

	return errs
}

func (client *ApiClient) SendConnectedClients(ctx context.Context, receiver *messages.Member, members *messages.ConnectedMembers) error {
	var errs error

	encoded, err := members.Encode()

	if err != nil {
		logger.Log.FailedToEncodeOutput(
			receiver.ConnectionId,
			receiver.ConnectionId,
			err,
		)

		errs = multierr.Append(errs, err)
	}

	input := &apigatewaymanagementapi.PostToConnectionInput{
		ConnectionId: aws.String(receiver.ConnectionId),
		Data:         encoded,
	}

	_, err = client.client.PostToConnection(ctx, input)

	if err != nil {
		logger.Log.FailedToSendMessage(
			receiver.ConnectionId,
			receiver.ConnectionId,
			err,
		)
		errs = multierr.Append(errs, err)
	}
	return errs
}
