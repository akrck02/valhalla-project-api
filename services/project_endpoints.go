package services

import (
	apimodels "github.com/akrck02/valhalla-core-sdk/models/api"
)

var CreateProjectEndpoint = apimodels.Endpoint{
	Path:     "",
	Method:   apimodels.PostMethod,
	Checks:   CreateProjectCheck,
	Listener: CreateProject,
	Secured:  true,
}

var GetProjectEndpoint = apimodels.Endpoint{
	Path:     "/{id}",
	Method:   apimodels.GetMethod,
	Checks:   GetProjectCheck,
	Listener: GetProject,
	Secured:  true,
}

var EditProjectEndpooint = apimodels.Endpoint{
	Path:     "",
	Method:   apimodels.PatchMethod,
	Checks:   EditProjectCheck,
	Listener: EditProject,
	Secured:  true,
}

var DeleteProjectEndpoint = apimodels.Endpoint{
	Path:     "/{id}",
	Method:   apimodels.DeleteMethod,
	Checks:   DeleteProjectCheck,
	Listener: DeleteProject,
	Secured:  true,
}
