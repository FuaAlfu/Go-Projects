package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

)

/*
using Go lambda function
*/

func HandleRequest(ctx context.Context, event *events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
	fmt.Println("event", event)

	return &events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       "You are awesome :)",
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}