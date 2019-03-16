package app

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{"response": "OK"})
	})

	// Health
	router.GET("/network", HealthController.GetNetworkInfo)

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

	//Project
	router.POST("/projects", ProjectController.AddProject)
	router.GET("/projects", ProjectController.GetProjects)
	router.GET("/projects/:id", ProjectController.GetProject)
	router.PUT("/projects/:id", ProjectController.UpdateProject)
	router.DELETE("/projects/:id", ProjectController.DeleteProject)

	//Task
	router.POST("/tasks", TaskController.AddTask)
	router.GET("/tasks", TaskController.GetTasks)
	router.GET("/tasks/:id", TaskController.GetTask)
	router.PUT("/tasks/:id", TaskController.UpdateTask)
	router.DELETE("/tasks/:id", TaskController.DeleteTask)

	return router
}
