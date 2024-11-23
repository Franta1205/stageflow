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
	MapRoutes()
	err := router.Run(":8080")
	if err != nil {
		defer initializers.CloseDB()
		panic(err)
	}
}
