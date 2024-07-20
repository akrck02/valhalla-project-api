package services

import (
	permissiondal "github.com/akrck02/valhalla-core-dal/services/permission"
	projectdal "github.com/akrck02/valhalla-core-dal/services/project"
	"github.com/akrck02/valhalla-core-sdk/http"
	projectmodels "github.com/akrck02/valhalla-core-sdk/models/project"
	systemmodels "github.com/akrck02/valhalla-core-sdk/models/system"
	"github.com/akrck02/valhalla-core-sdk/valerror"
	"github.com/gin-gonic/gin"
)

func CreateProject(context *systemmodels.ValhallaContext) (*systemmodels.Response, *systemmodels.Error) {

	project := context.Request.Body.(*projectmodels.Project)

	if project == nil {
		return nil, &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	perr := projectdal.CreateProject(context.Database.Client, project)
	if perr != nil {
		return nil, perr
	}

	return &systemmodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: "pong",
	}, nil
}

func DeleteProject(context *systemmodels.ValhallaContext) (*systemmodels.Response, *systemmodels.Error) {

	project := context.Request.Body.(*projectmodels.Project)

	// get if the user is the owner of the project to enable deletion
	canDelete := permissiondal.CanDeleteProject(context.Request.User, &projectmodels.Project{ID: project.ID})
	if !canDelete {
		return nil, &systemmodels.Error{
			Status:  http.HTTP_STATUS_FORBIDDEN,
			Error:   valerror.ACCESS_DENIED,
			Message: "Access denied: Cannot delete project",
		}
	}

	perr := projectdal.DeleteProject(context.Database.Client, project)
	if perr != nil {
		return nil, perr
	}

	return &systemmodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: gin.H{"message": "Project deleted"},
	}, nil
}

func EditProject(context *systemmodels.ValhallaContext) (*systemmodels.Response, *systemmodels.Error) {

	projectToEdit := context.Request.Body.(*projectmodels.Project)

	_, perr := projectdal.GetProject(context.Database.Client, projectToEdit)
	if perr != nil {
		return nil, &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.PROJECT_NOT_FOUND,
			Message: "Project not found",
		}
	}

	canEdit := permissiondal.CanEditProject(context.Request.User, projectToEdit)
	if !canEdit {
		return nil, &systemmodels.Error{
			Status:  http.HTTP_STATUS_FORBIDDEN,
			Error:   valerror.ACCESS_DENIED,
			Message: "Cannot edit user",
		}
	}

	updateErr := projectdal.EditProject(context.Database.Client, projectToEdit)
	if updateErr != nil {
		return nil, &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	return &systemmodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: systemmodels.Message{Message: "User updated"},
	}, nil
}

func GetProject(context *systemmodels.ValhallaContext) (*systemmodels.Response, *systemmodels.Error) {

	id := context.Request.Params.(*projectmodels.Project).ID

	project, perr := projectdal.GetProject(context.Database.Client, &projectmodels.Project{ID: id})
	if perr != nil {
		return nil, &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.PROJECT_NOT_FOUND,
			Message: "Project not found",
		}
	}

	return &systemmodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: project,
	}, nil
}
