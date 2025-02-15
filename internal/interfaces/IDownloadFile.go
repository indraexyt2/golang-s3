package interfaces

import (
	"context"
	"github.com/gin-gonic/gin"
)

type IDownloadService interface {
	DownloadFile(ctx context.Context, objectKey string) (string, error)
}

type IDownloadAPI interface {
	DownloadFile(c *gin.Context)
}
