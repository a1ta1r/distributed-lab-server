package main

import (
	"github.com/a1ta1r/distributed-lab-server/internal/app"
)

func main() {
	router := app.InitRouter()

	err := router.Run()

	if err != nil {
		panic(err)
	}
}
