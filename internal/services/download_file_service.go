package services

import (
	"context"
	"os"
	"s3-go-file-handling/internal/repositories"
)

type DownloadFileService struct {
	s3Repository *repositories.S3Repository
}

func NewDownloadFileService(s3Repository *repositories.S3Repository) *DownloadFileService {
	return &DownloadFileService{
		s3Repository: s3Repository,
	}
}

func (s *DownloadFileService) DownloadFile(ctx context.Context, objectKey string) (string, error) {
	return s.s3Repository.DownloadFile(ctx, os.Getenv("S3_BUCKET_NAME"), objectKey)
}
