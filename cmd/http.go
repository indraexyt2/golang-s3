package cmd

import (
	"github.com/gin-gonic/gin"
	"s3-go-file-handling/config"
	"s3-go-file-handling/helpers"
	"s3-go-file-handling/internal/api"
	"s3-go-file-handling/internal/repositories"
	"s3-go-file-handling/internal/services"
)

func SetupHTTP() {
	d := DependencyInjection()
	r := gin.Default()

	r.POST("/upload", d.UploadFileAPI.UploadFile)
	r.GET("/download", d.DownloadFileAPI.DownloadFile)

	err := r.Run(":8080")
	if err != nil {
		helpers.Logger.Error("Error running http server: ", err)
	}
}

type Dependencies struct {
	UploadFileAPI   *api.UploadAPI
	DownloadFileAPI *api.DownloadAPI
}

func DependencyInjection() *Dependencies {
	s3Repository := repositories.NewS3Repository(config.SetupS3Client())

	uploadSvc := services.NewUploadFileService(s3Repository)
	uploadApi := api.NewUploadAPI(uploadSvc)

	downloadSvc := services.NewDownloadFileService(s3Repository)
	downloadApi := api.NewDownloadAPI(downloadSvc)

	return &Dependencies{
		UploadFileAPI:   uploadApi,
		DownloadFileAPI: downloadApi,
	}
}
