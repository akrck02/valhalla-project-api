package services

import (
	"github.com/akrck02/valhalla-core-sdk/http"
	projectmodels "github.com/akrck02/valhalla-core-sdk/models/project"
	systemmodels "github.com/akrck02/valhalla-core-sdk/models/system"
	"github.com/akrck02/valhalla-core-sdk/valerror"
	"github.com/gin-gonic/gin"
)

var INVALID_REQUEST = &systemmodels.Error{
	Status:  http.HTTP_STATUS_BAD_REQUEST,
	Error:   valerror.INVALID_REQUEST,
	Message: "Invalid request",
}

func CreateProjectCheck(context *systemmodels.ValhallaContext, gin *gin.Context) *systemmodels.Error {
	project := &projectmodels.Project{}
	err := gin.ShouldBindJSON(project)
	if err != nil {
		return &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	context.Request.Body = project
	return nil
}

func DeleteProjectCheck(context *systemmodels.ValhallaContext, gin *gin.Context) *systemmodels.Error {

	project := &projectmodels.Project{}
	err := gin.ShouldBindJSON(project)
	if err != nil {
		return INVALID_REQUEST
	}

	context.Request.Body = project
	return nil
}

func EditProjectCheck(context *systemmodels.ValhallaContext, gin *gin.Context) *systemmodels.Error {

	project := &projectmodels.Project{}
	err := gin.ShouldBindJSON(project)
	if err != nil {
		return INVALID_REQUEST
	}

	context.Request.Body = project
	return nil
}

func GetProjectCheck(context *systemmodels.ValhallaContext, gin *gin.Context) *systemmodels.Error {

	return nil
}
