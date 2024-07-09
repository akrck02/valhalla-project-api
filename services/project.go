package services

import (
	server "github.com/akrck02/valhalla-api-common/http"
	permissiondal "github.com/akrck02/valhalla-core-dal/services/permission"
	projectdal "github.com/akrck02/valhalla-core-dal/services/project"
	"github.com/akrck02/valhalla-core-sdk/http"
	projectmodels "github.com/akrck02/valhalla-core-sdk/models/project"
	systemmodels "github.com/akrck02/valhalla-core-sdk/models/system"
	"github.com/akrck02/valhalla-core-sdk/valerror"
	"github.com/gin-gonic/gin"
)

func CreateProjectHttp(c *gin.Context) (*systemmodels.Response, *systemmodels.Error) {

	project := &projectmodels.Project{}
	err := c.ShouldBindJSON(project)
	if err != nil {
		return nil, &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	perr := projectdal.CreateProject(project)
	if perr != nil {
		return nil, perr
	}

	return &systemmodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: "pong",
	}, nil
}

func DeleteProjectHttp(c *gin.Context) (*systemmodels.Response, *systemmodels.Error) {

	request := server.GetRequestMetadata(c)
	project := &projectmodels.Project{}
	err := c.ShouldBindJSON(project)
	if err != nil {
		return nil, &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	// get if the user is the owner of the project to enable deletion
	canDelete := permissiondal.CanDeleteProject(request.User, &projectmodels.Project{ID: project.ID})
	if !canDelete {
		return nil, &systemmodels.Error{
			Status:  http.HTTP_STATUS_FORBIDDEN,
			Error:   valerror.ACCESS_DENIED,
			Message: "Access denied: Cannot delete project",
		}
	}

	perr := projectdal.DeleteProject(project)
	if perr != nil {
		return nil, perr
	}

	return &systemmodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: gin.H{"message": "Project deleted"},
	}, nil
}

// Edit user HTTP API endpoint
//
// [param] c | *gin.Context: context
func EditProjectHttp(c *gin.Context) (*systemmodels.Response, *systemmodels.Error) {

	//request := server.GetRequestMetadata(c)

	projectToEdit := &projectmodels.Project{}
	err := c.ShouldBindJSON(projectToEdit)
	if err != nil {
		return nil, &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}
	_, perr := projectdal.GetProject(projectToEdit)
	// get if request user is the owner of the project

	if perr != nil {
		return nil, &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.PROJECT_NOT_FOUND,
			Message: "Project not found",
		}
	}

	// canEdit := permissiondal.CanEditUser(request.User, userToEdit)
	// if !canEdit {
	// 	return nil, &models.Error{
	// 		Status:  http.HTTP_STATUS_FORBIDDEN,
	// 		Error:   error.ACCESS_DENIED,
	// 		Message: "Cannot edit user",
	// 	}
	// }

	updateErr := projectdal.EditProject(projectToEdit)
	if updateErr != nil {
		return nil, &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	return &systemmodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: gin.H{"message": "User updated"},
	}, nil
}
