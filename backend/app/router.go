package app

import (
	"fmt"
	"stageflow/controllers"
)

func MapRoutes() {
	fmt.Println("setting up routes")

	healthController := controllers.NewHealthController()
	router.GET("/ping", healthController.Ping)
}
