package app

import (
	"fmt"
	"stageflow/container"
	"stageflow/middlewares"
)

func MapRoutes(c *container.Container) {
	fmt.Println("setting up routes")
	router.Use(middlewares.SetUpCORS)
	router.GET("/ping", c.HealthController.Ping)

	public := router.Group("/api/v1")
	{
		auth := public.Group("/auth")
		{
			auth.POST("/register", c.AuthController.CreateUser)
			auth.POST("/login", c.AuthController.Login)
		}
	}

	protected := router.Group("/api/v1")
	protected.Use(middlewares.CheckAuth)
	{
		auth := protected.Group("/auth")
		{
			auth.GET("/user", c.AuthController.GetUser)
			auth.POST("/logout", c.AuthController.LogOut)
		}
		org := protected.Group("/organisation")
		{
			org.POST("/create", c.OrganisationController.Create)
			org.PUT("/:id/update", c.OrganisationController.Update)
		}
	}
}
