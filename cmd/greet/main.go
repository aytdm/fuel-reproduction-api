package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Result struct {
	Id int `json:"id"`
	Lang string `json:"lang"`
	Greeting string `json:"greeting"`
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	results := []Result{{
		Id:       1,
		Lang:     "en",
		Greeting: "Hello",
	}, {
		Id:       2,
		Lang:     "ja",
		Greeting: "こんにちは",
	}}
	results_byte, _ := json.Marshal(results)

	return &events.APIGatewayProxyResponse{
		Headers:           map[string]string{"Content-Type": "application/json; charset=UTF-8"},
		StatusCode:        200,
		Body:              string(results_byte),
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
