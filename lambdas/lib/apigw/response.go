package apigw

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type Response = events.APIGatewayProxyResponse

func InternalServerErrorResponse() Response {
	return Response{StatusCode: http.StatusInternalServerError, IsBase64Encoded: false}
}

func BadRequestResponse() Response {
	return Response{StatusCode: http.StatusBadRequest, IsBase64Encoded: false}
}

func OkResponse() Response {
	return Response{StatusCode: http.StatusOK, IsBase64Encoded: false}
}
