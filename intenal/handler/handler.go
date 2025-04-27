package handler

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/orgwats/config/intenal/service"
)

func Handler(ctx context.Context, req events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	log.Println("req.Body : ", req.Body)
	log.Println("req.Cookies : ", req.Cookies)
	log.Println("req.Headers : ", req.Headers)
	log.Println("req.QueryStringParameters : ", req.QueryStringParameters)
	log.Println("req.RawPath : ", req.RawPath)
	log.Println("req.RawQueryString : ", req.RawQueryString)
	log.Println("req.RequestContext : ", req.RequestContext)
	return service.GetConfig(req)
}
