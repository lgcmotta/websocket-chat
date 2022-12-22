package config

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/lgcmotta/websocket-chat/lib/logger"
	"go.uber.org/zap"
)

var Configuration aws.Config

func init() {
	var err error
	Configuration, err = config.LoadDefaultConfig(context.TODO())
	if err != nil {
		logger.Instance.Panic("unable to load SDK config",
			zap.Error(err),
		)
	}
}
