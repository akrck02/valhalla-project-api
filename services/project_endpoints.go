package services

import (
	"github.com/akrck02/valhalla-api-common/models"
	"github.com/akrck02/valhalla-core-sdk/http"
)

var CreateProjectEndpoint = models.Endpoint{
	Path:     "create",
	Method:   http.HTTP_METHOD_POST,
	Checks:   CreateProjectCheck,
	Listener: CreateProject,
	Database: true,
	Secured:  false,
}

var DeleteProjectEndpoint = models.Endpoint{
	Path:     "delete",
	Method:   http.HTTP_METHOD_DELETE,
	Checks:   DeleteProjectCheck,
	Listener: DeleteProject,
	Database: true,
	Secured:  false,
}

var EditProjectEndpooint = models.Endpoint{
	Path:     "edit",
	Method:   http.HTTP_METHOD_PUT,
	Checks:   EditProjectCheck,
	Listener: EditProject,
	Database: true,
	Secured:  false,
}

var GetProjectEndpoint = models.Endpoint{
	Path:     "get",
	Method:   http.HTTP_METHOD_GET,
	Checks:   GetProjectCheck,
	Listener: GetProject,
	Database: true,
	Secured:  false,
}
