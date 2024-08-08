package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"gitlab.com/blog/api/src/app/images"
	"gitlab.com/blog/api/src/config"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err.Error())
	}
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Test initialization of DynamoDB
	cfg, err := config.NewImages()
	if err != nil {
		return fmt.Errorf("load api config: %w", err)
	}

	awsConfig, err := awsconfig.LoadDefaultConfig(ctx)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	awsClient := s3.NewFromConfig(awsConfig)

	svc := images.New(cfg.S3.BucketName, awsClient)

	lambda.StartWithOptions(svc.GetUploadLink, lambda.WithContext(ctx))

	// Also change
	return nil
}
