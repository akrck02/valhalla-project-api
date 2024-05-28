package main

import (
	apicommon "github.com/akrck02/valhalla-api-common"
	"github.com/akrck02/valhalla-api-common/configuration"
	"github.com/akrck02/valhalla-api-common/models"
	databaseConfig "github.com/akrck02/valhalla-core-dal/configuration"
	"github.com/akrck02/valhalla-core-sdk/http"
	"github.com/akrck02/valhalla-project-api/services"
)

// main function
func main() {

	config := configuration.LoadConfiguration(".env")
	databaseConfig.LoadConfiguration(".env")

	apicommon.Start(
		config,
		[]models.Endpoint{

			// User endpoints
			{Path: "create", Method: http.HTTP_METHOD_POST, Listener: services.CreateProjectHttp},
			{Path: "delete", Method: http.HTTP_METHOD_DELETE, Listener: services.CreateProjectHttp},
			{Path: "edit", Method: http.HTTP_METHOD_PUT, Listener: services.CreateProjectHttp},
			{Path: "get", Method: http.HTTP_METHOD_GET, Listener: services.CreateProjectHttp},
		},
	)

}
