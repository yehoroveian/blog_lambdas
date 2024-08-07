package main

import (
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(HandleMessage)
}

func HandleMessage(message interface{}) (events.SNSEvent, error) {
	log.Printf("message: %+v", message)

	return events.SNSEvent{
		Records: []events.SNSEventRecord{
			{
				EventVersion:         "1",
				EventSubscriptionArn: "efqwe",
				EventSource:          "feqweq",
				SNS: events.SNSEntity{
					Signature:         "feqwe",
					MessageID:         "fqweqw",
					Type:              "fqweqw",
					TopicArn:          "qfweqw",
					MessageAttributes: nil,
					SignatureVersion:  "fwqeqw",
					Timestamp:         time.Now(),
					SigningCertURL:    "fqweqw",
					Message:           "fqweqw",
					UnsubscribeURL:    "fqweqw",
					Subject:           "fqwfq",
				},
			},
		},
	}, nil
}
