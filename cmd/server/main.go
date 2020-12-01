package main

import (
	"github.com/noworldwar/api-web/internal/app"
)

func main() {
	app.GenerateDoc()
	app.InitMySQL()
	app.InitRedis()
	app.InitRouter()
	go app.RunRouter()
	app.Cleanup()
}
