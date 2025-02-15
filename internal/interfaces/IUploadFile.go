package interfaces

import (
	"context"
	"github.com/gin-gonic/gin"
	"io"
)

type IUploadFileService interface {
	UploadFile(ctx context.Context, objectKey string, fileData io.Reader) error
}

type IUploadAPI interface {
	UploadFile(c *gin.Context)
}
