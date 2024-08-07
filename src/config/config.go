package config

import (
	"github.com/kelseyhightower/envconfig"
)

type S3 struct {
	BucketName string `envconfig:"ENV_S3_BUCKET_NAME" required:"true"`
}

type Images struct {
	S3 S3
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
