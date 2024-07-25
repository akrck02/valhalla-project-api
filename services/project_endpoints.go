package services

import (
	apimodels "github.com/akrck02/valhalla-core-sdk/models/api"
)

var CreateProjectEndpoint = apimodels.Endpoint{
	Path:             "",
	Method:           apimodels.PostMethod,
	Checks:           CreateProjectCheck,
	Listener:         CreateProject,
	Database:         true,
	Secured:          true,
	RequestMimeType:  apimodels.MimeApplicationJson,
	ResponseMimeType: apimodels.MimeApplicationJson,
}

var DeleteProjectEndpoint = apimodels.Endpoint{
	Path:             "",
	Method:           apimodels.DeleteMethod,
	Checks:           DeleteProjectCheck,
	Listener:         DeleteProject,
	Database:         true,
	Secured:          true,
	RequestMimeType:  apimodels.MimeApplicationJson,
	ResponseMimeType: apimodels.MimeApplicationJson,
}

var EditProjectEndpooint = apimodels.Endpoint{
	Path:             "",
	Method:           apimodels.PatchMethod,
	Checks:           EditProjectCheck,
	Listener:         EditProject,
	Database:         true,
	Secured:          true,
	RequestMimeType:  apimodels.MimeApplicationJson,
	ResponseMimeType: apimodels.MimeApplicationJson,
}

var GetProjectEndpoint = apimodels.Endpoint{
	Path:             "",
	Method:           apimodels.GetMethod,
	Checks:           GetProjectCheck,
	Listener:         GetProject,
	Database:         true,
	Secured:          true,
	RequestMimeType:  apimodels.MimeApplicationJson,
	ResponseMimeType: apimodels.MimeApplicationJson,
}
