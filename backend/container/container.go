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
	db := initializers.GetDB()
	// auth
	userRepository := repository.NewUserRepository(db)
	tokenRepository := repository.NewTokenRepository(initializers.GetRedisClient())
	authService := services.NewAuthService(userRepository, tokenRepository)

	// organisation
	organisationRepository := repository.NewOrganisationRepository(db)
	organisationService := services.NewOrganisationService(organisationRepository)

	return &Container{
		HealthController:       api.NewHealthController(),
		AuthController:         controllers.NewAuthenticationController(authService),
		OrganisationController: controllers.NewOrganisationController(organisationService),
	}
}
