package blog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Provider interface {
	AddBlogSection(ctx context.Context) error
}

type Service struct {
	bucketName string

	cl *s3.Client
}

// Interface compliance check.
var _ Provider = (*Service)(nil)

func New(bucketName string, awsClient *s3.Client) *Service {
	return &Service{
		bucketName: bucketName,

		cl: awsClient,
	}
}
