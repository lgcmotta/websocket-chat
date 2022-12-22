package main

import (
	"context"
	"sync"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/lgcmotta/websocket-chat/lib/apigw"
	"github.com/lgcmotta/websocket-chat/lib/logger"
	"github.com/lgcmotta/websocket-chat/lib/redis"
	"github.com/mediocregopher/radix/v4"
	"go.uber.org/zap"
)

type Stack struct {
	mu       sync.Mutex
	elements []string
}

func (s *Stack) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.elements)
}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, req *events.APIGatewayWebsocketProxyRequest) (apigw.Response, error) {
	defer func() {
		_ = logger.Instance.Sync()
	}()

	logger.Instance.Info("websocket disconnecting inactive connections",
		zap.String("requestId", req.RequestContext.RequestID),
		zap.String("connectionId", req.RequestContext.ConnectionID))

	stack := new(Stack)
	redis.Client.Do(ctx, radix.Cmd(&(stack.elements), "SMEMBERS", "connections"))

	logger.Instance.Info("websocket connections read from cache",
		zap.Int("connections", stack.Len()),
		zap.String("requestId", req.RequestContext.RequestID),
		zap.String("connectionId", req.RequestContext.ConnectionID))

	var result string

	for _, connectionId := range stack.elements {
		if connectionId == req.RequestContext.ConnectionID {
			continue
		}

		err := redis.Client.Do(ctx, radix.Cmd(&result, "SREM", "connections", connectionId))
		if err != nil {
			logger.Instance.Error("failed to delete connection details from cache",
				zap.String("requestId", req.RequestContext.RequestID),
				zap.String("connectionId", req.RequestContext.ConnectionID),
				zap.Error(err))

			return apigw.InternalServerErrorResponse(), err
		}
	}

	logger.Instance.Info("websocket connections deleted from cache",
		zap.String("requestId", req.RequestContext.RequestID),
		zap.String("connectionId", req.RequestContext.ConnectionID))

	return apigw.OkResponse(), nil
}
