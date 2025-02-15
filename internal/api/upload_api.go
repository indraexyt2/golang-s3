package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"s3-go-file-handling/helpers"
	"s3-go-file-handling/internal/interfaces"
	"strings"
)

type UploadAPI struct {
	uploadFileService interfaces.IUploadFileService
}

func NewUploadAPI(uploadFileService interfaces.IUploadFileService) *UploadAPI {
	return &UploadAPI{
		uploadFileService: uploadFileService,
	}
}

func (api *UploadAPI) UploadFile(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	file, err := c.FormFile("file")
	if err != nil {
		log.Info("Error uploading file: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, "Error uploading file", nil)
		return
	}

	allowedExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".webp": true,
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedExtensions[ext] {
		log.Info("Invalid file type: ", ext)
		helpers.SendResponse(c, http.StatusBadRequest, "Only image files are allowed", nil)
		return
	}

	fileData, err := file.Open()
	if err != nil {
		log.Info("Error opening file: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, "Error opening file", nil)
		return
	}
	defer fileData.Close()

	objectKey := fmt.Sprintf("uploads/%s", file.Filename)
	err = api.uploadFileService.UploadFile(c.Request.Context(), objectKey, fileData)
	if err != nil {
		log.Info("Error uploading file: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, "Error uploading file", nil)
		return
	}

	helpers.SendResponse(c, http.StatusOK, "File uploaded successfully", nil)
}
