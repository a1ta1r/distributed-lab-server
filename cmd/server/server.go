package main

import (
	"github.com/a1ta1r/distributed-lab-server/internal/app"
	"os"
)

func main() {
	app.UpdateSchema()

	router := app.InitRouter()

	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "3001"
	}

	println("Server listening on port " + port)
	err := router.Run(":" + port)

	if err != nil {
		panic(err)
	}
}
