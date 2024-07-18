package services

import (
	"github.com/akrck02/valhalla-api-common/models"
	"github.com/akrck02/valhalla-core-sdk/http"
	"github.com/akrck02/valhalla-project-api/services"
)

var CreateProjectEndpoint = models.Endpoint{
	Path:     "create",
	Method:   http.HTTP_METHOD_POST,
	Checks:   services.CreateProjectCheck,
	Listener: services.CreateProject,
	Database: true,
}

var DeleteProjectEndpoint = models.Endpoint{
	Path:     "delete",
	Method:   http.HTTP_METHOD_DELETE,
	Checks:   services.DeleteProjectCheck,
	Listener: services.DeleteProject,
	Database: true,
}

var EditProjectEndpooint = models.Endpoint{
	Path:     "edit",
	Method:   http.HTTP_METHOD_PUT,
	Checks:   services.EditProjectCheck,
	Listener: services.EditProject,
	Database: true,
}

var GetProjectEndpoint = models.Endpoint{
	Path:     "get",
	Method:   http.HTTP_METHOD_GET,
	Checks:   services.GetProjectCheck,
	Listener: services.GetProject,
	Database: true,
}
