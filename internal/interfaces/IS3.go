package interfaces

import (
	"context"
	"io"
)

type IS3Repository interface {
	UploadFile(ctx context.Context, bucketName string, objectKey string, fileData io.Reader) error
	DownloadFile(ctx context.Context, bucketName string, objectKey string) (string, error)
}
