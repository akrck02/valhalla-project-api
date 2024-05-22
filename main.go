package main

import (
	apicommon "github.com/akrck02/valhalla-api-common"
	"github.com/akrck02/valhalla-api-common/configuration"
	"github.com/akrck02/valhalla-api-common/models"
	"github.com/akrck02/valhalla-core-sdk/http"
	"github.com/akrck02/valhalla-project-api/services"
)

// main function
func main() {

	config := configuration.LoadConfiguration(".env")

	apicommon.Start(
		config,
		[]models.Endpoint{

			// User endpoints
			{Path: "ping", Method: http.HTTP_METHOD_GET, Listener: services.Ping},
		},
	)

}
