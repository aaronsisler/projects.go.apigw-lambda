package handlers

import (
	"context"

	"github.com/aaronsisler/projects.go.apigw-lambda/shared"
	"github.com/aws/aws-lambda-go/events"
)

func HelloGetHandler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	name := req.QueryStringParameters["name"]

	message := shared.FormatMessage(name)

	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       "Handler: Hello: GET: " + message,
	}, nil
}
