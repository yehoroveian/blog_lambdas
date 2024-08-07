package images

import (
	"context"

	"gitlab.com/blog/api/src/storage/dynamo/namespaces"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Provider interface {
	GetUploadLink(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}

type Service struct {
	bucketName string

	db namespaces.Repository
	cl *s3.Client
}

// Interface compliance check.
var _ Provider = (*Service)(nil)

func New(bucketName string, awsClient *s3.Client, repository namespaces.Repository) *Service {
	return &Service{
		bucketName: bucketName,

		db: repository,
		cl: awsClient,
	}
}
