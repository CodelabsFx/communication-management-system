package storage

import (
    "context"
    "net/http"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/credentials"
    "github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewS3Client(ctx context.Context, endpoint, accessKey, secretKey string) (*s3.Client, error) {
    cfg, err := config.LoadDefaultConfig(ctx,
        config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
        config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
            return aws.Endpoint{URL: endpoint, SigningRegion: "us-east-1", HostnameImmutable: true}, nil
        })),
    )
    if err != nil {
        return nil, err
    }
    return s3.NewFromConfig(cfg, func(o *s3.Options) {
        o.UsePathStyle = true
    }), nil
}
