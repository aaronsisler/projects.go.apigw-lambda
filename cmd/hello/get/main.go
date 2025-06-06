package main

import (
	"github.com/aaronsisler/projects.go.apigw-lambda/handlers/hello"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(hello.HelloGetHandler)
}
