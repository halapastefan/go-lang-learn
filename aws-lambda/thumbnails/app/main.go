package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)
type Event events.S3EventRecord

func Handler(event Event) (string, error) {

	fmt.Print(event)
	fmt.Print("something1")
	return "I am lambda", nil
}

func main() {
	lambda.Start(Handler)
}
