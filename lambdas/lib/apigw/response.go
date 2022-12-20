package apigw

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type Response = events.APIGatewayProxyResponse

func InternalServerErrorResponse() Response {
	return Response{StatusCode: http.StatusInternalServerError}
}

func BadRequestResponse() Response {
	return Response{StatusCode: http.StatusBadRequest}
}

func OkResponse() Response {
	return Response{StatusCode: http.StatusOK}
}
