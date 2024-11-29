package app

import (
	"github.com/gin-gonic/gin"
	"stageflow/config/initializers"
)

var (
	router = gin.Default()
)

func StartApp() {
	initializers.ConnectDB()
	initializers.ConnectRedis()
	MapRoutes()
	defer initializers.CloseDB()
	defer initializers.CloseRedis()
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
