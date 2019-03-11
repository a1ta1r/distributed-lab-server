package app

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{"response": "OK"})
	})

	// User
	router.POST("/users", UserController.AddUser)
	router.GET("/users", UserController.GetUsers)
	router.GET("/users/:id", UserController.GetUser)
	router.PUT("/users/:id", UserController.UpdateUser)
	router.DELETE("/users/:id", UserController.DeleteUser)

	// Team
	router.POST("/teams", TeamController.AddTeam)
	router.GET("/teams", TeamController.GetTeams)
	router.GET("/teams/:id", TeamController.GetTeam)
	router.PUT("/teams/:id", TeamController.UpdateTeam)
	router.DELETE("/teams/:id", TeamController.DeleteTeam)

	//Status
	router.POST("/statuses", StatusController.AddStatus)
	router.GET("/statuses", StatusController.GetStatuses)
	router.GET("/statuses/:id", StatusController.GetStatus)
	router.PUT("/statuses/:id", StatusController.UpdateStatus)
	router.DELETE("/statuses/:id", StatusController.DeleteStatus)

	return router
}
