package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"s3-go-file-handling/helpers"
	"s3-go-file-handling/internal/services"
)

type UploadAPI struct {
	uploadFileService *services.UploadFileService
}

func NewUploadAPI(uploadFileService *services.UploadFileService) *UploadAPI {
	return &UploadAPI{
		uploadFileService: uploadFileService,
	}
}

func (api *UploadAPI) UploadFile(c *gin.Context) {
	var (
		log = helpers.Logger
		ctx = c.Request.Context()
	)

	file, err := c.FormFile("file")
	if err != nil {
		log.Info("Error uploading file: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, "Error uploading file", nil)
		return
	}

	fileData, err := file.Open()
	if err != nil {
		log.Info("Error opening file: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, "Error opening file", nil)
		return
	}
	defer fileData.Close()

	objectKey := "uploads/" + file.Filename
	err = api.uploadFileService.UploadFile(ctx, objectKey, fileData)
	if err != nil {
		log.Info("Error uploading file: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, "Error uploading file", nil)
		return
	}

	helpers.SendResponse(c, http.StatusOK, "File uploaded successfully", nil)
}
