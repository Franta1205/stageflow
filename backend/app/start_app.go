package app

import (
	"github.com/gin-gonic/gin"
	"stageflow/config/initializers"
	"stageflow/container"
)

var (
	router = gin.Default()
)

func StartApp() {
	initializers.ConnectDB()
	initializers.ConnectRedis()
	// to init controllers
	c := container.NewContainer()
	MapRoutes(c)
	defer initializers.CloseDB()
	defer initializers.CloseRedis()
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
