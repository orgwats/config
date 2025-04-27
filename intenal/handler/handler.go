package handler

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/orgwats/config/intenal/service"
)

func Handler(ctx context.Context, req events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	if req.RawPath != "/config" {
		return events.LambdaFunctionURLResponse{StatusCode: 404, Body: "Not Found"}, nil
	} else {
		return service.GetConfig(req)
	}
}
