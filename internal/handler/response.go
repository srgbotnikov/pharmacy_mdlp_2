package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"error" xml:"error"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Println(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{Message: message})
}

type statusResponse struct {
	Status string `json:"status" xml:"status"`
}
