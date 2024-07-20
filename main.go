package main

import (
	apicommon "github.com/akrck02/valhalla-api-common"
	"github.com/akrck02/valhalla-api-common/configuration"
	"github.com/akrck02/valhalla-api-common/models"
	databaseConfig "github.com/akrck02/valhalla-core-dal/configuration"

	"github.com/akrck02/valhalla-project-api/services"
)

const ENVIROMENT_FILE = ".env"

// main function
func main() {

	config := configuration.LoadConfiguration(ENVIROMENT_FILE)
	databaseConfig.LoadConfiguration(ENVIROMENT_FILE)

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
