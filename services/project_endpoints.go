package services

import (
	"github.com/akrck02/valhalla-core-sdk/http"
	apimodels "github.com/akrck02/valhalla-core-sdk/models/api"
)

var CreateProjectEndpoint = apimodels.Endpoint{
	Path:     "",
	Method:   http.HTTP_METHOD_POST,
	Checks:   CreateProjectCheck,
	Listener: CreateProject,
	Database: true,
	Secured:  true,
}

var DeleteProjectEndpoint = apimodels.Endpoint{
	Path:     "",
	Method:   http.HTTP_METHOD_DELETE,
	Checks:   DeleteProjectCheck,
	Listener: DeleteProject,
	Database: true,
	Secured:  true,
}

var EditProjectEndpooint = apimodels.Endpoint{
	Path:     "",
	Method:   http.HTTP_METHOD_PATCH,
	Checks:   EditProjectCheck,
	Listener: EditProject,
	Database: true,
	Secured:  true,
}

var GetProjectEndpoint = apimodels.Endpoint{
	Path:     "",
	Method:   http.HTTP_METHOD_GET,
	Checks:   GetProjectCheck,
	Listener: GetProject,
	Database: true,
	Secured:  true,
}
