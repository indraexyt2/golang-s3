package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"s3-go-file-handling/helpers"
	"s3-go-file-handling/internal/services"
)

type DownloadAPI struct {
	downloadFileService *services.DownloadFileService
}

func NewDownloadAPI(downloadFileService *services.DownloadFileService) *DownloadAPI {
	return &DownloadAPI{
		downloadFileService: downloadFileService,
	}
}

func (api *DownloadAPI) DownloadFile(c *gin.Context) {
	fileName := c.Query("file")
	objectKey := "uploads/" + fileName

	if fileName == "" {
		helpers.SendResponse(c, http.StatusBadRequest, "File name is required", nil)
		return
	}

	url, err := api.downloadFileService.DownloadFile(c.Request.Context(), objectKey)
	if err != nil {
		helpers.SendResponse(c, http.StatusInternalServerError, "Error downloading file", nil)
		return
	}

	helpers.SendResponse(c, http.StatusOK, "File downloaded successfully", url)
}
