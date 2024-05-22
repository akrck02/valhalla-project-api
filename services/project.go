package services

import (
	"github.com/akrck02/valhalla-core-sdk/http"
	"github.com/akrck02/valhalla-core-sdk/models"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) (*models.Response, *models.Error) {
	return &models.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: "pong",
	}, nil
}
