package services

import (
	permissiondal "github.com/akrck02/valhalla-core-dal/services/permission"
	projectdal "github.com/akrck02/valhalla-core-dal/services/project"
	"github.com/akrck02/valhalla-core-sdk/http"
	apimodels "github.com/akrck02/valhalla-core-sdk/models/api"
	projectmodels "github.com/akrck02/valhalla-core-sdk/models/project"
	"github.com/akrck02/valhalla-core-sdk/valerror"
)

func CreateProject(context *apimodels.ApiContext) (*apimodels.Response, *apimodels.Error) {

	project := context.Request.Body.(*projectmodels.Project)

	if project == nil {
		return nil, &apimodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	perr := projectdal.CreateProject(context.Database.Client, project)
	if perr != nil {
		return nil, perr
	}

	return &apimodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: "pong",
	}, nil
}

func DeleteProject(context *apimodels.ApiContext) (*apimodels.Response, *apimodels.Error) {

	project := context.Request.Body.(*projectmodels.Project)

	// get if the user is the owner of the project to enable deletion
	canDelete := permissiondal.CanDeleteProject(context.Trazability.User, &projectmodels.Project{ID: project.ID})
	if !canDelete {
		return nil, &apimodels.Error{
			Status:  http.HTTP_STATUS_FORBIDDEN,
			Error:   valerror.ACCESS_DENIED,
			Message: "Access denied: Cannot delete project",
		}
	}

	perr := projectdal.DeleteProject(context.Database.Client, project)
	if perr != nil {
		return nil, perr
	}

	return &apimodels.Response{
		Code: http.HTTP_STATUS_OK,
		Response: apimodels.Message{
			Message: "Project deleted",
		},
	}, nil
}

func EditProject(context *apimodels.ApiContext) (*apimodels.Response, *apimodels.Error) {

	projectToEdit := context.Request.Body.(*projectmodels.Project)

	_, perr := projectdal.GetProject(context.Database.Client, projectToEdit)
	if perr != nil {
		return nil, &apimodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.PROJECT_NOT_FOUND,
			Message: "Project not found",
		}
	}

	canEdit := permissiondal.CanEditProject(context.Trazability.User, projectToEdit)
	if !canEdit {
		return nil, &apimodels.Error{
			Status:  http.HTTP_STATUS_FORBIDDEN,
			Error:   valerror.ACCESS_DENIED,
			Message: "Cannot edit user",
		}
	}

	updateErr := projectdal.EditProject(context.Database.Client, projectToEdit)
	if updateErr != nil {
		return nil, &apimodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	return &apimodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: apimodels.Message{Message: "User updated"},
	}, nil
}

func GetProject(context *apimodels.ApiContext) (*apimodels.Response, *apimodels.Error) {

	id := context.Request.Params.(*projectmodels.Project).ID

	project, perr := projectdal.GetProject(context.Database.Client, &projectmodels.Project{ID: id})
	if perr != nil {
		return nil, &apimodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.PROJECT_NOT_FOUND,
			Message: "Project not found",
		}
	}

	return &apimodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: project,
	}, nil
}
