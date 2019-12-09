package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandlerFunc(c context.Context) (string, error) {
	return "I am lambda 1", nil
}

func main() {

	lambda.Start(HandlerFunc)
}
