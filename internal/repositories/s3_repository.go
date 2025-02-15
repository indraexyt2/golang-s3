package repositories

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
	"time"
)

type S3Repository struct {
	s3client *s3.Client
}

func NewS3Repository(s3client *s3.Client) *S3Repository {
	return &S3Repository{
		s3client: s3client,
	}
}

func (r *S3Repository) UploadFile(ctx context.Context, bucketName string, objectKey string, fileData io.Reader) error {
	_, err := r.s3client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
		Body:   fileData,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *S3Repository) DownloadFile(ctx context.Context, bucketName string, objectKey string) (string, error) {
	presignedClient := s3.NewPresignClient(r.s3client)
	presignedUrl, err := presignedClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
	}, s3.WithPresignExpires(time.Minute*10))

	if err != nil {
		return "", err
	}

	return presignedUrl.URL, nil
}
