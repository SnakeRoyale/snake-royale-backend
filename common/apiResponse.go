package common

import (
	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Message string            `json:"message"`
}

func WriteFailApiResponse(c *gin.Context, httpCode int, message string) {
	c.AbortWithStatusJSON(httpCode, ApiResponse{
		Message: message,
	})
}

func WriteOKApiResponse(c *gin.Context, httpCode int, message string) {
	c.JSON(httpCode, ApiResponse{
		Message: message,
	})
}