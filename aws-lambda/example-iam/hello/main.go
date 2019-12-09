package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context) (string, error) {
	return "I am lamda", nil
}

func main() {
	lambda.Start(Handler)
}
