package handler

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/orgwats/config/intenal/service"
)

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("req.HTTPMethod : ", req.HTTPMethod)
	log.Println("req.Headers : ", req.Headers)
	log.Println("req.Body : ", req.Body)
	log.Println("req.IsBase64Encoded : ", req.IsBase64Encoded)
	log.Println("req.MultiValueHeaders : ", req.MultiValueHeaders)
	log.Println("req.MultiValueQueryStringParameters : ", req.MultiValueQueryStringParameters)
	log.Println("req.Path : ", req.Path)
	log.Println("req.PathParameters : ", req.PathParameters)
	log.Println("req.QueryStringParameters : ", req.QueryStringParameters)
	log.Println("req.RequestContext : ", req.RequestContext)
	log.Println("req.Resource : ", req.Resource)
	log.Println("req.StageVariables : ", req.StageVariables)
	log.Println("req.RequestContext : ", req.RequestContext)
	return service.GetConfig(req)
}
