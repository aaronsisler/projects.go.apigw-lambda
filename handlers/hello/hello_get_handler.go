package hello

import (
	"context"
	"encoding/json"

	"github.com/aaronsisler/projects.go.apigw-lambda/shared"
	"github.com/aws/aws-lambda-go/events"
)

func HelloGetHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := req.QueryStringParameters["name"]

	message := shared.FormatMessage(name)

	responseBody, err := json.Marshal(map[string]string{
		"message": "Handler: Hello: GET: " + message,
	})

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Internal Server Error"}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(responseBody),
	}, nil
}
