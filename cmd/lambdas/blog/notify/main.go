package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(PrintMessage)
}

func PrintMessage() {
	log.Printf("SUCCESS TO RUN NOTIFY BLOG")
}
