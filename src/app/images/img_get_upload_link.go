package images

import (
	"context"
	"encoding/json"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

func (s *Service) GetUploadLink(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	client := s3.NewPresignClient(s.cl, s3.WithPresignExpires(10*time.Minute))
	key := uuid.New().String() + ".jpg"

	req, err := client.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: &s.bucketName,
		Key:    &key,
	})
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	resp := struct {
		UploadURL string `json:"upload_url"`
		Key       string `json:"key"`
	}{
		UploadURL: req.URL,
		Key:       key,
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode:      200,
		Body:            string(jsonResp),
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}
