package services

import (
	"context"
	"io"
	"os"
	"s3-go-file-handling/internal/interfaces"
)

type UploadFileService struct {
	s3Repository interfaces.IS3Repository
}

func NewUploadFileService(s3Repository interfaces.IS3Repository) *UploadFileService {
	return &UploadFileService{
		s3Repository: s3Repository,
	}
}

func (s *UploadFileService) UploadFile(ctx context.Context, objectKey string, fileData io.Reader) error {
	return s.s3Repository.UploadFile(ctx, os.Getenv("S3_BUCKET_NAME"), objectKey, fileData)
}
