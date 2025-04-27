package service

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client = s3.New(session.Must(session.NewSession()))
	bucket   = os.Getenv("BUCKET")
)

func GetConfig(req events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	serviceParam := req.QueryStringParameters["service"]
	if serviceParam == "" {
		return errorResponse(400, "Missing 'service' query parameter")
	}

	serviceNames := strings.Split(serviceParam, ",")

	result, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String("config.json"),
	})
	if err != nil {
		log.Println("S3 error:", err)
		return errorResponse(404, "Config file not found")
	}

	var config map[string]interface{}
	if err := json.NewDecoder(result.Body).Decode(&config); err != nil {
		log.Println("Decode error:", err)
		return errorResponse(500, "Failed to decode config file")
	}

	commonConfig, ok := config["common"]
	if !ok {
		return errorResponse(500, "Missing 'common' in config file")
	}

	serviceConfigs, ok := config["service"].(map[string]interface{})
	if !ok {
		return errorResponse(500, "Missing 'service' section in config file")
	}

	selectedServices := make(map[string]interface{})
	for _, name := range serviceNames {
		name = strings.TrimSpace(name)
		if svc, ok := serviceConfigs[name]; ok {
			selectedServices[name] = svc
		} else {
			return errorResponse(400, fmt.Sprintf("Unknown service: %s", name))
		}
	}

	merged := map[string]interface{}{
		"common":  commonConfig,
		"service": selectedServices,
	}

	body, _ := json.MarshalIndent(merged, "", "  ")
	return events.LambdaFunctionURLResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(body),
	}, nil
}

func errorResponse(status int, message string) (events.LambdaFunctionURLResponse, error) {
	errorBody, _ := json.Marshal(map[string]string{
		"error": message,
	})
	return events.LambdaFunctionURLResponse{
		StatusCode: status,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(errorBody),
	}, nil
}
