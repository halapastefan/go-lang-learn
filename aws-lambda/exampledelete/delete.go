package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleLambda() (string, error) {
	return "deleting", nil
}
func main() {
	lambda.Start(HandleLambda)
}
