package logger

import (
	"github.com/aws/aws-lambda-go/events"
	"go.uber.org/zap"
)

type Logger struct {
	instance *zap.Logger
}

var Log *Logger

func init() {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}

	logger, _ := config.Build()

	Log = &Logger{
		instance: logger,
	}
}

func (logs *Logger) Sync() {
	_ = logs.instance.Sync()
}

func (logs *Logger) RequestStarted(req *events.APIGatewayWebsocketProxyRequest) {
	logs.instance.Info("websocket request started",
		zap.String("routeKey", req.RequestContext.RouteKey),
		zap.String("requestId", req.RequestContext.RequestID),
		zap.String("connectionId", req.RequestContext.ConnectionID),
	)
}

func (logs *Logger) RequestEnded(req *events.APIGatewayWebsocketProxyRequest) {
	logs.instance.Info("websocket request ended",
		zap.String("routeKey", req.RequestContext.RouteKey),
		zap.String("requestId", req.RequestContext.RequestID),
		zap.String("connectionId", req.RequestContext.ConnectionID),
	)
	logs.Sync()
}

func (logs *Logger) ConnectionIdSaved(req *events.APIGatewayWebsocketProxyRequest) {
	logs.instance.Info("websocket connection id saved",
		zap.String("requestId", req.RequestContext.RequestID),
		zap.String("connectionId", req.RequestContext.ConnectionID),
	)
}

func (logs *Logger) FailedToSaveConnectionId(req *events.APIGatewayWebsocketProxyRequest, err error) {
	logs.instance.Error("failed to save connection id",
		zap.String("requestId", req.RequestContext.RequestID),
		zap.String("connectionId", req.RequestContext.ConnectionID),
		zap.Error(err),
	)
}
func (logs *Logger) FailedToDeleteConnectionId(req *events.APIGatewayWebsocketProxyRequest, err error) {
	logs.instance.Error("failed to delete connection id",
		zap.String("requestId", req.RequestContext.RequestID),
		zap.String("connectionId", req.RequestContext.ConnectionID),
		zap.Error(err),
	)
}

func (logs *Logger) FailedToUpdateMemberNickname(req *events.APIGatewayWebsocketProxyRequest, nickname string, err error) {
	logs.instance.Error("failed to update member nickname",
		zap.String("requestId", req.RequestContext.RequestID),
		zap.String("connectionId", req.RequestContext.ConnectionID),
		zap.String("nickname", nickname),
		zap.Error(err),
	)
}

func (logs *Logger) FailedToDecodeInput(req *events.APIGatewayWebsocketProxyRequest, err error) {
	logs.instance.Error("failed to parse client input",
		zap.String("requestId", req.RequestContext.RequestID),
		zap.String("connectionId", req.RequestContext.ConnectionID),
		zap.String("input", req.Body),
		zap.Error(err),
	)
}

func (logs *Logger) FailedToRetrieveMembers(req *events.APIGatewayWebsocketProxyRequest, err error) {
	logs.instance.Error("failed to retrive connected chat members",
		zap.String("requestId", req.RequestContext.RequestID),
		zap.String("connectionId", req.RequestContext.ConnectionID),
		zap.Error(err),
	)
}

func (logs *Logger) FailedToRetrieveSender(req *events.APIGatewayWebsocketProxyRequest, err error) {
	logs.instance.Error("failed to retrieve sender member",
		zap.String("requestId", req.RequestContext.RequestID),
		zap.String("connectionId", req.RequestContext.ConnectionID),
		zap.Error(err),
	)
}

func (logs *Logger) FailedToRetrieveReceiver(req *events.APIGatewayWebsocketProxyRequest, err error) {
	logs.instance.Error("failed to retrieve receiver member",
		zap.String("requestId", req.RequestContext.RequestID),
		zap.String("connectionId", req.RequestContext.ConnectionID),
		zap.Error(err),
	)
}

func (logs *Logger) FailedToRetrieveLeavingMember(req *events.APIGatewayWebsocketProxyRequest, err error) {
	logs.instance.Error("failed to retrieve leaving member",
		zap.String("requestId", req.RequestContext.RequestID),
		zap.String("connectionId", req.RequestContext.ConnectionID),
		zap.Error(err),
	)
}

func (logs *Logger) FailedToEncodeOutput(senderId, receiverId string, err error) {
	logs.instance.Error("failed to encode message output",
		zap.String("senderId", senderId),
		zap.String("receiverId", receiverId),
		zap.Error(err),
	)
}

func (logs *Logger) FailedToSendMessage(senderId, receiverId string, err error) {
	logs.instance.Error("send message failed",
		zap.String("senderId", senderId),
		zap.String("receiverId", receiverId),
		zap.Error(err),
	)
}

func (logs *Logger) FailedToMarshalItemDynamoDB(err error) {
	logs.instance.Error("failed to marshal dynamodb item",
		zap.Error(err),
	)
}

func (logs *Logger) FailedToUnmarshalItemDynamoDB(err error) {
	logs.instance.Error("failed to unmarshal dynamodb item(s)",
		zap.Error(err),
	)
}

func (logs *Logger) FailedToPutItemDynamoDB(partitionKey string, err error) {
	logs.instance.Error("failed to put item on dynamodb",
		zap.String("partitionId", partitionKey),
		zap.Error(err),
	)
}

func (logs *Logger) FailedToScanItemsDynamoDB(err error) {
	logs.instance.Error("failed to scan items from dynamodb",
		zap.Error(err),
	)
}

func (logs *Logger) FailedToGetItemDynamoDB(partitionKey string, err error) {
	logs.instance.Error("failed to get item from dynamodb",
		zap.String("partitionId", partitionKey),
		zap.Error(err),
	)
}

func (logs *Logger) FailedToDeleteItemDynamoDB(partitionKey string, err error) {
	logs.instance.Error("failed to delete item from dynamodb",
		zap.String("partitionId", partitionKey),
		zap.Error(err),
	)
}

func (logs *Logger) PanicLoadingSDKConfiguration(err error) {
	logs.instance.Panic("panic: unable to load AWS configuration",
		zap.Error(err),
	)
}
