package services

import (
	projectdal "github.com/akrck02/valhalla-core-dal/services/project"
	"github.com/akrck02/valhalla-core-sdk/http"
	"github.com/akrck02/valhalla-core-sdk/models"
	"github.com/akrck02/valhalla-core-sdk/valerror"
	"github.com/gin-gonic/gin"
)

func CreateProjectHttp(c *gin.Context) (*models.Response, *models.Error) {

	project := &models.Project{}
	err := c.ShouldBindJSON(project)
	if err != nil {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	perr := projectdal.CreateProject(project)
	if perr != nil {
		return nil, perr
	}

	return &models.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: "pong",
	}, nil
}

// Edit user HTTP API endpoint
//
// [param] c | *gin.Context: context
func EditProjectHttp(c *gin.Context) (*models.Response, *models.Error) {

	//request := server.GetRequestMetadata(c)

	projectToEdit := &models.Project{}
	err := c.ShouldBindJSON(projectToEdit)
	if err != nil {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}
	_, perr := projectdal.GetProject(projectToEdit)
	// get if request user is the owner of the project

	if perr != nil {
		return nil, &models.Error{
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
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	return &models.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: gin.H{"message": "User updated"},
	}, nil
}
