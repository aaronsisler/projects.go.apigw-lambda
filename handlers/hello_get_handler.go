package handlers

import (
	"context"

	"github.com/aaronsisler/projects.go.apigw-lambda/shared"
	"github.com/aws/aws-lambda-go/events"
)

func HelloGetHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := req.QueryStringParameters["name"]

	message := shared.FormatMessage(name)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: "Handler: Hello: GET: " + message,
	}, nil
}
