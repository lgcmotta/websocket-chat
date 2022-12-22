package apigw

import (
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/lgcmotta/websocket-chat/lib/logger"
	"go.uber.org/zap"
)

func NewAPIGatewayManagementClient(cfg *aws.Config, domain, stage string) *apigatewaymanagementapi.Client {

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		logger.Instance.Info("inputs from endpoint resolver", zap.String("service", service), zap.String("region", region))

		if service != "execute-api" {
			return cfg.EndpointResolverWithOptions.ResolveEndpoint(service, region, options...)
		}

		var endpoint url.URL
		endpoint.Path = stage
		endpoint.Host = domain
		endpoint.Scheme = "https"
		return aws.Endpoint{
			SigningRegion: region,
			URL:           endpoint.String(),
		}, nil
	})

	cfg.EndpointResolverWithOptions = customResolver

	return apigatewaymanagementapi.NewFromConfig(*cfg)
}
