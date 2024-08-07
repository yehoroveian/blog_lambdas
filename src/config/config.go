package config

import (
	"github.com/kelseyhightower/envconfig"
)

type ECR struct {
	URI  string `envconfig:"ENV_ECR_URI" required:"true"`
	User string `envconfig:"ENV_ECR_IMAGE_TAG" required:"true"`
}

type Secret struct {
	SecretName   string `envconfig:"ENV_SECRET_NAME" required:"true"`
	SecretRegion string `envconfig:"ENV_SECRET_REGION" required:"true"`
}

type S3 struct {
	BucketName string `envconfig:"ENV_S3_BUCKET_NAME" required:"true"`
}

type Lambda struct {
	LambdaRunnerRole string `envconfig:"ENV_LAMBDA_RUNNER_ROLE" required:"true"`
}

type Deploy struct {
	ECR    ECR
	Secret Secret
	Lambda Lambda
}

type Images struct {
	S3 S3
}

func NewDeploy() (Deploy, error) {
	return load[Deploy]()
}

func NewImages() (Images, error) {
	return load[Images]()
}

func load[T any]() (T, error) {
	const prefix = ""

	var config T
	if err := envconfig.Process(prefix, &config); err != nil {
		return config, err
	}

	return config, nil
}
