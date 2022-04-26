package main

import(
	"fmt"
	"os"

	"github.com/aws/aws-lambada-go/events"
	_"github.com/aws/aws-lambada-go/lambada"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const tableName = "lambdaInGoUser"

func handler(req events.APIGetewayProxyRequest)(*events.APIGetewayProxyRequest, error){
	switch req.HTTPMethod {
		case "GET":
			return handlers.GetUser(req, tableName, dynaClient)
		case "POST":
			return handlers.CreateUser(req,tableName, dynaClient)
		case "PUT":
			return handlers.UpdateUser(req, tableName, dynaClient)
		case "DELETE":
			return handlers.UnhandledMethod()
	}
}

func main() {
	region := os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region)
	},)
	if err != nil{
		return nil
	}
	dynaClient = dynamodb.New(awsSession)
	lambda.Start(handler)
}