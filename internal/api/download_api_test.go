package api

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"s3-go-file-handling/helpers"
	"testing"
)

type MockDownloadFileService struct {
	mock.Mock
}

func (m *MockDownloadFileService) DownloadFile(ctx context.Context, objectKey string) (string, error) {
	args := m.Called(ctx, objectKey)
	if args.Error(1) != nil {
		return "", args.Error(1)
	}
	return args.String(0), nil
}

func setupRouterDownloadFile(downloadSvc *MockDownloadFileService) *gin.Engine {
	r := gin.Default()
	api := NewDownloadAPI(downloadSvc)
	r.GET("/download", api.DownloadFile)
	return r
}

func TestDownloadAPI_DownloadFile(t *testing.T) {
	helpers.Logger = &MockLogger{}

	t.Run("can download file", func(t *testing.T) {
		mockDownloadFileSvc := new(MockDownloadFileService)
		mockDownloadFileSvc.On("DownloadFile", mock.Anything, mock.Anything).Return("https://example.com/test.txt", nil)

		r := setupRouterDownloadFile(mockDownloadFileSvc)

		req, err := http.NewRequest("GET", "/download?file=test.txt", nil)
		assert.Nil(t, err)

		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		var resp helpers.Response
		err = json.Unmarshal(rec.Body.Bytes(), &resp)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "File downloaded successfully", resp.Message)
		assert.NotNil(t, resp.Data)

		mockDownloadFileSvc.AssertExpectations(t)
	})

	t.Run("can't download file", func(t *testing.T) {
		mockDownloadFileSvc := new(MockDownloadFileService)
		mockDownloadFileSvc.On("DownloadFile", mock.Anything, mock.Anything).Return(nil, errors.New("error downloading file"))

		r := setupRouterDownloadFile(mockDownloadFileSvc)

		req, err := http.NewRequest("GET", "/download?file=test.txt", nil)
		assert.Nil(t, err)

		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		var resp helpers.Response
		err = json.Unmarshal(rec.Body.Bytes(), &resp)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, "Error downloading file", resp.Message)
		assert.Nil(t, resp.Data)

		mockDownloadFileSvc.AssertExpectations(t)
	})
}
