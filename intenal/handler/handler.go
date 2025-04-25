package handler

import (
	"context"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/orgwats/config/intenal/service"
)

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch {
	case strings.HasPrefix(req.Path, "/config"):
		return service.GetConfig(req)
	default:
		return events.APIGatewayProxyResponse{
			StatusCode: 404,
			Body:       "Not Found",
		}, nil
	}
}
