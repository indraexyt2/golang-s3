package helpers

import "github.com/gin-gonic/gin"

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SendResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, Response{
		Message: message,
		Data:    data,
	})
}
