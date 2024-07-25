package services

import (
	"net/http"

	projectdal "github.com/akrck02/valhalla-core-dal/services/project"
	apierror "github.com/akrck02/valhalla-core-sdk/error"
	apimodels "github.com/akrck02/valhalla-core-sdk/models/api"
	projectmodels "github.com/akrck02/valhalla-core-sdk/models/project"
)

func CreateProject(context *apimodels.ApiContext) (*apimodels.Response, *apimodels.Error) {

	project := context.Request.Body.(*projectmodels.Project)
	project, perr := projectdal.CreateProject(context.Database.Client, project)
	if perr != nil {
		return nil, perr
	}

	return &apimodels.Response{
		Code:     http.StatusOK,
		Response: project,
	}, nil
}

func DeleteProject(context *apimodels.ApiContext) (*apimodels.Response, *apimodels.Error) {

	project := context.Request.Body.(*projectmodels.Project)
	perr := projectdal.DeleteProject(context.Database.Client, project)
	if perr != nil {
		return nil, perr
	}

	return &apimodels.Response{
		Code: http.StatusOK,
		Response: apimodels.Message{
			Message: "Project deleted",
		},
	}, nil
}

func EditProject(context *apimodels.ApiContext) (*apimodels.Response, *apimodels.Error) {

	project := context.Request.Body.(*projectmodels.Project)
	updateErr := projectdal.EditProject(context.Database.Client, project)
	if updateErr != nil {
		return nil, &apimodels.Error{
			Status:  http.StatusBadRequest,
			Error:   apierror.InvalidRequest,
			Message: "Invalid request",
		}
	}

	project, perr := projectdal.GetProject(context.Database.Client, project)
	if perr != nil {
		return nil, &apimodels.Error{
			Status:  http.StatusInternalServerError,
			Error:   apierror.ProjectNotFound,
			Message: "Project not found",
		}
	}

	return &apimodels.Response{
		Code:     http.StatusOK,
		Response: project,
	}, nil
}

func GetProject(context *apimodels.ApiContext) (*apimodels.Response, *apimodels.Error) {

	id := context.Request.Params["id"]

	project, perr := projectdal.GetProject(context.Database.Client, &projectmodels.Project{ID: id})
	if perr != nil {
		return nil, &apimodels.Error{
			Status:  http.StatusBadRequest,
			Error:   apierror.ProjectNotFound,
			Message: "Project not found",
		}
	}

	return &apimodels.Response{
		Code:     http.StatusOK,
		Response: project,
	}, nil
}
