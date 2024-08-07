package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type OutEvent struct {
	ImageID      string `json:"image_id"`
	ResizeFormat string `json:"resize_format"`
}

func main() {
	lambda.Start(HandleMessage)
}

func HandleMessage(message events.SQSMessage) (OutEvent, error) {
	log.Printf("handling message: %+v", message)

	return OutEvent{
		ImageID:      "Hello World!",
		ResizeFormat: "Save",
	}, nil
}
