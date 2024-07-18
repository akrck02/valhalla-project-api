package main

import (
	apicommon "github.com/akrck02/valhalla-api-common"
	"github.com/akrck02/valhalla-api-common/configuration"
	"github.com/akrck02/valhalla-api-common/models"
	databaseConfig "github.com/akrck02/valhalla-core-dal/configuration"

	"github.com/akrck02/valhalla-project-api/services"
)

// main function
func main() {

	config := configuration.LoadConfiguration(".env")
	databaseConfig.LoadConfiguration(".env")

	apicommon.Start(
		config,
		[]models.Endpoint{
			services.CreateProjectEndpoint,
			services.DeleteProjectEndpoint,
			services.EditProjectEndpooint,
			services.GetProjectEndpoint,
		},
	)

}
