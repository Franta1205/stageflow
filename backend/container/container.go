package container

import (
	"stageflow/api"
	"stageflow/api/v1/controllers"
	"stageflow/api/v1/repository"
	"stageflow/api/v1/services"
	"stageflow/config/initializers"
)

type Container struct {
	HealthController       *api.HealthController
	AuthController         *controllers.AuthController
	OrganisationController *controllers.OrganisationController
}

func NewContainer() *Container {
	// auth
	userRepository := repository.NewUserRepository()
	tokenRepository := repository.NewTokenRepository(initializers.GetRedisClient())
	authService := services.NewAuthService(userRepository, tokenRepository)

	// organisation
	organisationService := services.NewOrganisationService()

	return &Container{
		HealthController:       api.NewHealthController(),
		AuthController:         controllers.NewAuthenticationController(authService),
		OrganisationController: controllers.NewOrganisationController(organisationService),
	}
}