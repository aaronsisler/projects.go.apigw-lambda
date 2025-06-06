package handlers

import (
	"context"
	"encoding/json"

	"github.com/aaronsisler/projects.go.apigw-lambda/shared"
	"github.com/aws/aws-lambda-go/events"
)

type HelloRequest struct {
	Name string `json:"name"`
}

func HelloPostHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var helloRequest HelloRequest

	// Unmarshal the JSON string in the body
	err := json.Unmarshal([]byte(req.Body), &helloRequest)

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Invalid request body"}, nil
	}

	message := shared.FormatMessage(helloRequest.Name)

	responseBody, err := json.Marshal(map[string]string{
		"message": "Handler: Hello: POST: " + message,
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
