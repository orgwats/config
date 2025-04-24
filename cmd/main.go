package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/orgwats/config/intenal/handler"
)

func main() {
	lambda.Start(handler.Handler)
}
