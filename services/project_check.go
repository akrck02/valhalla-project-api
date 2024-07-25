package services

import (
	"io"
	"net/http"

	permissiondal "github.com/akrck02/valhalla-core-dal/services/permission"
	projectdal "github.com/akrck02/valhalla-core-dal/services/project"
	apierror "github.com/akrck02/valhalla-core-sdk/error"
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

	id := context.Request.Params["id"]

	if id == "" {
		return INVALID_REQUEST
	}

	project, perr := projectdal.GetProject(context.Database.Client, &projectmodels.Project{ID: id})
	if perr != nil {
		return &apimodels.Error{
			Status:  http.StatusNotFound,
			Error:   apierror.ProjectNotFound,
			Message: "Project not found",
		}
	}

	// get if the user is the owner of the project to enable deletion
	canDelete := permissiondal.CanDeleteProject(context.Trazability.User, project)
	if !canDelete {
		return &apimodels.Error{
			Status:  http.StatusForbidden,
			Error:   apierror.AccessDenied,
			Message: "Access denied: Cannot delete project",
		}
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

	databaseProject, perr := projectdal.GetProject(context.Database.Client, project)
	if perr != nil {
		return &apimodels.Error{
			Status:  http.StatusBadRequest,
			Error:   apierror.ProjectNotFound,
			Message: "Project not found",
		}
	}

	canEdit := permissiondal.CanEditProject(context.Trazability.User, databaseProject)
	if !canEdit {
		return &apimodels.Error{
			Status:  http.StatusForbidden,
			Error:   apierror.AccessDenied,
			Message: "Access denied: Cannot edit project",
		}
	}

	context.Request.Body = project
	return nil
}

func GetProjectCheck(context *apimodels.ApiContext) *apimodels.Error {

	id := context.Request.Params["id"]

	if id == "" {
		return INVALID_REQUEST
	}

	project, perr := projectdal.GetProject(context.Database.Client, &projectmodels.Project{ID: id})
	if perr != nil {
		return &apimodels.Error{
			Status:  http.StatusNotFound,
			Error:   apierror.ProjectNotFound,
			Message: "Project not found",
		}
	}

	canView := permissiondal.CanSeeProject(context.Trazability.User, project)
	if !canView {
		return &apimodels.Error{
			Status:  http.StatusForbidden,
			Error:   apierror.AccessDenied,
			Message: "Access denied: Cannot view project",
		}
	}

	return nil
}
