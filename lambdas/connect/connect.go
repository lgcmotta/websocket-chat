package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/lgcmotta/websocket-chat/lib/apigw"
)

func main() {
	lambda.Start(handler)
}

func handler(_ context.Context, req *events.APIGatewayWebsocketProxyRequest) (apigw.Response, error) {

	return apigw.OkResponse(), nil
}
