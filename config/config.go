package config

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
	"os"
	"s3-go-file-handling/helpers"
)

func SetupConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		helpers.Logger.Error("Error loading .env file")
		return
	}
	helpers.Logger.Info("Config loaded")
}

func SetupS3Client() *s3.Client {
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(os.Getenv("AWS_REGION")))
	if err != nil {
		helpers.Logger.Error("Error loading AWS config: ", err)
		return nil
	}

	s3client := s3.NewFromConfig(cfg)
	return s3client
}
