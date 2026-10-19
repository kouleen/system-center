package utils

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	fileconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/kouleen/idl/kitex_gen/system"
)

var fileClient *s3.Client
var bucket string

func init() {
	accessKeyID := os.Getenv("CLOUDFLARE_ACCESS_KEY_ID")
	if accessKeyID == "" {
		log.Fatalf("CLOUDFLARE_ACCESS_KEY_ID environment variable not set")
	}
	secretAccessKey := os.Getenv("CLOUDFLARE_SECRET_ACCESS_KEY")
	if secretAccessKey == "" {
		log.Fatalf("CLOUDFLARE_SECRET_ACCESS_KEY environment variable not set")
	}
	endpoint := os.Getenv("CLOUDFLARE_ENDPOINT")
	if endpoint == "" {
		log.Fatalf("CLOUDFLARE_ENDPOINT environment variable not set")
	}
	b := os.Getenv("CLOUDFLARE_BUCKET")
	if b == "" {
		log.Fatalf("CLOUDFLARE_BUCKET environment variable not set")
	}
	cfg, err := fileconfig.LoadDefaultConfig(context.TODO(),
		fileconfig.WithRegion("auto"),
		fileconfig.WithCredentialsProvider(aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
			return aws.Credentials{AccessKeyID: accessKeyID, SecretAccessKey: secretAccessKey}, nil
		})),
	)
	if err != nil {
		log.Fatalf("LoadDefaultConfig failed: %v", err)
	}
	cli := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endpoint)
	})
	fileClient = cli
	bucket = b
}

func UploadFile(ctx context.Context, req *system.SystemFileRequest) (resp *system.SystemFileResponse, err error) {
	now := time.Now()
	year := now.Year()        // int 年份，如 2026
	month := int(now.Month()) // time.Month 类型（本质int，1=1月，12=12月）
	Key := fmt.Sprintf("%d/%d/%s", year, month, req.GetFileName())
	input := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(Key),
		Body:   bytes.NewReader(req.GetFileData()),
		// 建议带上ContentLength和ContentType
		ContentLength: aws.Int64(req.GetFileSize()),
		ContentType:   aws.String(req.GetContentType()),
	}
	_, err = fileClient.PutObject(ctx, input)
	if err != nil {
		return nil, err
	}
	return &system.SystemFileResponse{
		FileName:    req.GetFileName(),
		FileSize:    req.GetFileSize(),
		ContentType: req.GetContentType(),
		DownloadUrl: "https://file.kouleen.cn/" + Key,
	}, nil
}
