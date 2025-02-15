package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"s3-go-file-handling/helpers"
	"testing"
)

type MockUploadFileService struct {
	mock.Mock
}

func (m *MockUploadFileService) UploadFile(ctx context.Context, objectKey string, fileData io.Reader) error {
	args := m.Called(ctx, objectKey, fileData)
	return args.Error(0)
}

type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Info(msg string, args ...interface{})  {}
func (m *MockLogger) Error(msg string, args ...interface{}) {}

func setupRouter(uploadSvc *MockUploadFileService) *gin.Engine {
	r := gin.Default()
	api := NewUploadAPI(uploadSvc)
	r.POST("/upload", api.UploadFile)
	return r
}

func TestUploadAPI_UploadFile(t *testing.T) {
	helpers.Logger = &MockLogger{}

	t.Run("can upload file", func(t *testing.T) {
		mockUploadFileSvc := new(MockUploadFileService)
		mockUploadFileSvc.On("UploadFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)

		r := setupRouter(mockUploadFileSvc)

		buffer := &bytes.Buffer{}
		writer := multipart.NewWriter(buffer)

		file, err := writer.CreateFormFile("file", "test.txt")
		assert.Nil(t, err)

		_, err = io.Copy(file, bytes.NewBuffer([]byte("test content")))
		assert.Nil(t, err)

		err = writer.Close()
		assert.Nil(t, err)

		req, err := http.NewRequest("POST", "/upload", buffer)
		assert.Nil(t, err)

		req.Header.Set("Content-Type", writer.FormDataContentType())
		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		var resp helpers.Response
		err = json.Unmarshal(rec.Body.Bytes(), &resp)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "File uploaded successfully", resp.Message)

		mockUploadFileSvc.AssertExpectations(t)
	})

	t.Run("can't upload file", func(t *testing.T) {
		mockUploadFileSvc := new(MockUploadFileService)
		mockUploadFileSvc.On("UploadFile", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error uploading file"))

		r := setupRouter(mockUploadFileSvc)

		buffer := &bytes.Buffer{}
		writer := multipart.NewWriter(buffer)

		file, err := writer.CreateFormFile("file", "test.txt")
		assert.Nil(t, err)

		_, err = io.Copy(file, bytes.NewBuffer([]byte("test content")))
		assert.Nil(t, err)

		err = writer.Close()
		assert.Nil(t, err)

		req, err := http.NewRequest("POST", "/upload", buffer)
		assert.Nil(t, err)

		req.Header.Set("Content-Type", writer.FormDataContentType())
		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		var resp helpers.Response
		err = json.Unmarshal(rec.Body.Bytes(), &resp)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, "Error uploading file", resp.Message)

		mockUploadFileSvc.AssertExpectations(t)
	})

}
