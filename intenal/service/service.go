package service

import (
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client = s3.New(session.Must(session.NewSession()))
	bucket   = os.Getenv("BUCKET")
)

func GetConfig(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	result, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String("config.json"),
	})
	if err != nil {
		log.Println("S3 error:", err)
		return events.APIGatewayProxyResponse{StatusCode: 404, Body: "Config not found"}, nil
	}

	var config map[string]interface{}
	if err := json.NewDecoder(result.Body).Decode(&config); err != nil {
		log.Println("Decode error:", err)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Decode error"}, nil
	}

	body, _ := json.MarshalIndent(config, "", "  ")
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(body),
	}, nil
}
