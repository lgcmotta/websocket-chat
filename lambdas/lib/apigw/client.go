package apigw

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/lgcmotta/websocket-chat/lib/logger"
	"go.uber.org/zap"
)

func NewAPIGatewayManagementClient(cfg *aws.Config, domain, stage string) *apigatewaymanagementapi.Client {

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		logger.Instance.Info("inputs from endpoint resolver", zap.String("service", service), zap.String("region", region))

		return aws.Endpoint{
			SigningRegion: region,
			URL:           "https://0pltkpqrxl.execute-api.us-east-1.amazonaws.com/production/",
		}, nil
	})

	cfg.EndpointResolverWithOptions = customResolver

	return apigatewaymanagementapi.NewFromConfig(*cfg)
}
