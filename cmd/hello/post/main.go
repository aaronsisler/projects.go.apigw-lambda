package main

import (
	"github.com/aaronsisler/projects.go.apigw-lambda/handlers"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handlers.HelloPostHandler)
}
