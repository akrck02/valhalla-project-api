package services

import (
	"github.com/akrck02/valhalla-core-sdk/http"
	apimodels "github.com/akrck02/valhalla-core-sdk/models/api"
	projectmodels "github.com/akrck02/valhalla-core-sdk/models/project"
	"github.com/akrck02/valhalla-core-sdk/valerror"
)

var INVALID_REQUEST = &apimodels.Error{
	Status:  http.HTTP_STATUS_BAD_REQUEST,
	Error:   valerror.INVALID_REQUEST,
	Message: "Invalid request",
}

func CreateProjectCheck(context *apimodels.ApiContext) *apimodels.Error {
	project := &projectmodels.Project{}
	// err := gin.ShouldBindJSON(project)
	// if err != nil {
	// 	return &apimodels.Error{
	// 		Status:  http.HTTP_STATUS_BAD_REQUEST,
	// 		Error:   valerror.INVALID_REQUEST,
	// 		Message: "Invalid request",
	// 	}
	// }

	context.Request.Body = project
	return nil
}

func DeleteProjectCheck(context *apimodels.ApiContext) *apimodels.Error {

	project := &projectmodels.Project{}
	// err := gin.ShouldBindJSON(project)
	// if err != nil {
	// 	return INVALID_REQUEST
	// }

	context.Request.Body = project
	return nil
}

func EditProjectCheck(context *apimodels.ApiContext) *apimodels.Error {

	project := &projectmodels.Project{}
	// err := gin.ShouldBindJSON(project)
	// if err != nil {
	// 	return INVALID_REQUEST
	// }

	context.Request.Body = project
	return nil
}

func GetProjectCheck(context *apimodels.ApiContext) *apimodels.Error {

	return nil
}
