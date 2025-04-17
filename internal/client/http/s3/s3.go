package s3

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Config struct {
	Secret_key string `envconfig:"SECRET_KEY"`
	Access_key string `envconfig:"ACCESS_KEY"`
	Bucketname string `envconfig:"BUCKETNAME"`
	Region     string `envconfig:"REGION"`
	Endpoint   string `envconfig:"ENDPOINT"`
}

func Connect(s3Cfg *S3Config) *s3.Client {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(s3Cfg.Region),
		config.WithBaseEndpoint(s3Cfg.Endpoint),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(s3Cfg.Access_key, s3Cfg.Secret_key, "")),
	)
	if err != nil {
		panic(err)
	}
	s3Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	return s3Client
}
