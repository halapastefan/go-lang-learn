package main

import "github.com/aws/aws-lambda-go/lambda"

type CompareData struct {
	Age    int `json:"name"`
	Height int `json:"height"`
	Income int `json:"income"`
}

func HandleLambda(event CompareData) (CompareData, error) {
	return event, nil
}
func main() {

	lambda.Start(HandleLambda)
}
