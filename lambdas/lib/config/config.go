package config

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/lgcmotta/websocket-chat/lib/logger"
)

var Configuration aws.Config

func init() {
	var err error
	Configuration, err = config.LoadDefaultConfig(context.TODO())
	if err != nil {
		logger.Log.PanicLoadingSDKConfiguration(err)
		panic(err)
	}
}
