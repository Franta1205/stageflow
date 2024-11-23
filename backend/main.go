package main

import (
	"fmt"
	"stageflow/app"
	"stageflow/config/initializers"
)

func main() {
	f := initializers.LoadEnvVariable("DB_USER")
	fmt.Println(f)
	app.StartApp()
}
