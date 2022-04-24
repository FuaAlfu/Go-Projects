package main

import(
	"fmt"
	"os"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	region := os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region)
	},)
}