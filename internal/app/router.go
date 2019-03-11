package app

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{"response": "OK"})
	})

	router.POST("/users", UserController.AddUser)
	router.GET("/users", UserController.GetUsers)
	router.GET("/users/:id", UserController.GetUser)
	router.PUT("/users/:id", UserController.UpdateUser)
	router.DELETE("/users/:id", UserController.DeleteUser)

	return router
}
