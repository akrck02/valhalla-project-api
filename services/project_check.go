package services

import (
	"io"
	"net/http"

	apierror "github.com/akrck02/valhalla-core-sdk/error"
	"github.com/akrck02/valhalla-core-sdk/log"
	apimodels "github.com/akrck02/valhalla-core-sdk/models/api"
	projectmodels "github.com/akrck02/valhalla-core-sdk/models/project"
	"github.com/akrck02/valhalla-core-sdk/utils"
)

var INVALID_REQUEST = &apimodels.Error{
	Status:  http.StatusBadRequest,
	Error:   apierror.InvalidRequest,
	Message: "Invalid request",
}

func CreateProjectCheck(context *apimodels.ApiContext) *apimodels.Error {
	project := &projectmodels.Project{}
	err := utils.ParseJson(context.Request.Body.(io.Reader), project)
	if err != nil {
		return INVALID_REQUEST
	}

	context.Request.Body = project
	return nil
}

func DeleteProjectCheck(context *apimodels.ApiContext) *apimodels.Error {

	project := &projectmodels.Project{}
	err := utils.ParseJson(context.Request.Body.(io.Reader), project)
	if err != nil {
		return INVALID_REQUEST
	}

	context.Request.Body = project
	return nil
}

func EditProjectCheck(context *apimodels.ApiContext) *apimodels.Error {

	project := &projectmodels.Project{}
	err := utils.ParseJson(context.Request.Body.(io.Reader), project)
	if err != nil {
		return INVALID_REQUEST
	}

	context.Request.Body = project
	return nil
}

func GetProjectCheck(context *apimodels.ApiContext) *apimodels.Error {

	urlValues := &map[string]string{}

	err := utils.ParseJson(context.Request.Body.(io.Reader), urlValues)

	if err != nil {
		log.Error(err.Error())
		return INVALID_REQUEST
	}

	context.Request.Body = &projectmodels.Project{
		ID: (*urlValues)["id"],
	}

	return nil
}
