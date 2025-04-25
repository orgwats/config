package handler

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/orgwats/config/intenal/service"
)

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println(req)
	return service.GetConfig(req)
}
