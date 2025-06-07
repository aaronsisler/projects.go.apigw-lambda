package main

import (
	"github.com/aaronsisler/projects.go.apigw-lambda/handlers/user"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(user.UserGetHandler)
}
