package handler

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/orgwats/config/intenal/service"
)

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return service.GetConfig(req)
}
