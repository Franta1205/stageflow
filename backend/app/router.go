package app

import (
	"fmt"
	"stageflow/api"
	"stageflow/api/v1/controllers"
	"stageflow/middlewares"
)

func MapRoutes() {
	fmt.Println("setting up routes")

	healthController := api.NewHealthController()
	router.GET("/ping", healthController.Ping)

	public := router.Group("/api/v1")
	{
		authController := controllers.NewAuthenticationController()
		auth := public.Group("/auth")
		{
			auth.POST("/register", authController.CreateUser)
			auth.POST("/login", authController.Login)
		}
	}

	protected := router.Group("/api/v1")
	protected.Use(middlewares.CheckAuth)
	{
		authController := controllers.NewAuthenticationController()
		auth := protected.Group("/auth")
		{
			auth.GET("/user", middlewares.CheckAuth, authController.GetUser)
		}
	}
}
