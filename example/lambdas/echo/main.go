package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	b, err := json.MarshalIndent(request, "", "    ")
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(b),
	}, err
}

func main() {
	lambda.Start(Handler)
}
