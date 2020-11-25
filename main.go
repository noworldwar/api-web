package main

import (
	"github.com/noworldwar/api-web/app"
)

func main() {
	app.InitMySQL()
	app.InitRedis()
	app.InitRouter()
	go app.RunRouter()
	app.Cleanup()
}
