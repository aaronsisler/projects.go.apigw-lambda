package user

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aaronsisler/projects.go.apigw-lambda/dao"
	"github.com/aaronsisler/projects.go.apigw-lambda/service"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var userService *service.UserService

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("failed to load config:", err)
		os.Exit(1)
	}

	ddb := dynamodb.NewFromConfig(cfg)
	userDao := dao.NewUserDao(ddb, "SERVICES_EVENTS_ADMIN_SERVICE")
	userService = service.NewUserService(userDao)
}

func UserGetHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	userId := req.PathParameters["userId"]
	if userId == "" {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Missing userId"}, nil
	}

	user, err := userService.GetUserById(ctx, userId)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: err.Error()}, nil
	}

	if user == nil {
		return events.APIGatewayProxyResponse{StatusCode: 404, Body: "User not found"}, nil
	}

	body, _ := json.Marshal(user)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(body),
	}, nil
}
