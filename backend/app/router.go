package app

import (
	"fmt"
	"stageflow/api"
	"stageflow/api/v1/controllers"
)

func MapRoutes() {
	fmt.Println("setting up routes")

	healthController := api.NewHealthController()
	router.GET("/ping", healthController.Ping)

	v1Api := router.Group("/api/v1")
	{
		authController := controllers.NewAuthenticationController()
		auth := v1Api.Group("/auth")
		{
			auth.POST("/register", authController.CreateUser)
		}
	}
}
