package app

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{"response": "OK"})
	})

	router.GET("/users", UserController.GetUsers)

	return router
}
